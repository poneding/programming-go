package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.etcd.io/etcd/clientv3"
)

type Register struct {
	EtcdEndpoints []string
	DialTimeout   int
	closeCh       chan struct{}
	leaseID       clientv3.LeaseID
	keepAliveCh   <-chan *clientv3.LeaseKeepAliveResponse

	server    Server
	serverTTL int64
	cli       *clientv3.Client
	logger    *logrus.Logger
}

// NewRegister 可以考虑使用函数式选项模式
func NewRegister(etcdEndpoints []string, logger *logrus.Logger) *Register {
	return &Register{
		EtcdEndpoints: etcdEndpoints,
		DialTimeout:   3,
		closeCh:       make(chan struct{}),
	}
}

func (r *Register) Register(server Server, ttl int64) (chan<- struct{}, error) {
	var err error
	if strings.Split(server.Address, ":")[0] == "" {
		return nil, errors.New("invalid ip address")
	}
	if r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdEndpoints,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	}); err != nil {
		return nil, err
	}

	r.server = server
	r.serverTTL = ttl
	if r.register(); err != nil {
		return nil, err
	}

	r.closeCh = make(chan struct{})
	go r.keepAlive()

	return r.closeCh, nil
}

func (r *Register) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeout)*time.Second)
	defer cancel()

	leaseResp, err := r.cli.Grant(ctx, r.serverTTL)
	if err != nil {
		return err
	}

	r.leaseID = leaseResp.ID

	if r.keepAliveCh, err = r.cli.KeepAlive(context.Background(), r.leaseID); err != nil {
		return err
	}

	data, err := json.Marshal(r.server)
	if err != nil {
		return err
	}
	_, err = r.cli.Put(context.Background(), BuildRegisterPath(r.server), string(data), clientv3.WithLease(r.leaseID))

	return err
}

func (r *Register) Stop() {
	r.closeCh <- struct{}{}
}

func (r *Register) unregister() error {
	_, err := r.cli.Delete(context.Background(), BuildRegisterPath(r.server))
	return err
}

func (r *Register) keepAlive() {
	ticker := time.NewTicker(time.Duration(r.serverTTL) * time.Second)
	for {
		select {
		case <-r.closeCh:
			if err := r.unregister(); err != nil {
				r.logger.Errorf("unregister failed, err: %s", err.Error())
			}
			if _, err := r.cli.Revoke(context.Background(), r.leaseID); err != nil {
				r.logger.Errorf("revoke failed, err: %s", err.Error())
			}
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					r.logger.Errorf("register failed, err: %s", err.Error())
				}
			}
		case <-ticker.C:
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					r.logger.Errorf("register failed, err: %s", err.Error())
				}
			}
		}
	}
}

func (r *Register) UpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		weight, err := cast.ToInt64E(req.URL.Query().Get("weight"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		update := func() error {
			r.server.Weight = weight
			data, err := json.Marshal(r.server)
			if err != nil {
				return err
			}
			_, err = r.cli.Put(context.Background(), BuildRegisterPath(r.server), string(data), clientv3.WithLease(r.leaseID))
			return err
		}

		if err := update(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("update server weight success"))
	}
}

func (r *Register) GetServer() (Server, error) {
	resp, err := r.cli.Get(context.Background(), BuildRegisterPath(r.server))
	if err != nil {
		return r.server, err
	}

	var server Server
	if resp.Count > 0 {
		if err := json.Unmarshal(resp.Kvs[0].Value, &server); err != nil {
			return server, err
		}
	}
	return server, nil
}
