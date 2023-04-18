package discovery

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
)

const (
	schema = "etcd"
)

type Resolver struct {
	schema        string
	EtcdEndpoints []string
	DialTimeout   int

	closeCh         chan struct{}
	watchCh         clientv3.WatchChan
	cli             *clientv3.Client
	keyPrefix       string
	serverAddresses []resolver.Address

	conn   resolver.ClientConn
	logger *logrus.Logger
}

func NewResolver(etcdEndpoints []string, logger *logrus.Logger) *Resolver {
	return &Resolver{
		schema:        schema,
		EtcdEndpoints: etcdEndpoints,
		DialTimeout:   3,
		logger:        logger,
	}
}

func (r *Resolver) Scheme() string {
	return r.schema
}

func (r *Resolver) Build(target resolver.Target, conn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.conn = conn
	r.keyPrefix = BuildPrefix(Server{Name: target.URL.Path, Version: target.URL.Host})
	if _, err := r.start(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Resolver) ResolveNow(opts resolver.ResolveNowOptions) {}

func (r *Resolver) Close() {
	r.closeCh <- struct{}{}
}

func (r *Resolver) start() (chan<- struct{}, error) {
	var err error
	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdEndpoints,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}

	resolver.Register(r)

	r.closeCh = make(chan struct{})

	if err := r.sync(); err != nil {
		return nil, err
	}
	go r.watch()

	return r.closeCh, nil
}

func (r *Resolver) update(events []*clientv3.Event) {
	for _, ev := range events {
		switch ev.Type {
		case clientv3.EventTypePut:
			s, err := ParseValue(ev.Kv.Value)
			if err != nil {
				continue
			}

			addr := resolver.Address{Addr: s.Address, Metadata: s.Weight}
			if !Exist(r.serverAddresses, addr) {
				r.serverAddresses = append(r.serverAddresses, addr)
				r.conn.UpdateState(resolver.State{Addresses: r.serverAddresses})
			}
		case clientv3.EventTypeDelete:
			s, err := SplitPath(string(ev.Kv.Key))
			if err != nil {
				continue
			}
			addr := resolver.Address{Addr: s.Address}
			if s, ok := Remove(r.serverAddresses, addr); ok {
				r.serverAddresses = s
				r.conn.UpdateState(resolver.State{Addresses: r.serverAddresses})
			}
		}
	}
}

func (r *Resolver) sync() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := r.cli.Get(ctx, r.keyPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	r.serverAddresses = []resolver.Address{}

	for _, v := range res.Kvs {
		info, err := ParseValue(v.Value)
		if err != nil {
			continue
		}
		addr := resolver.Address{Addr: info.Address, Metadata: info.Weight}
		r.serverAddresses = append(r.serverAddresses, addr)
	}

	r.conn.UpdateState(resolver.State{Addresses: r.serverAddresses})
	return nil
}

func (r *Resolver) watch() {
	ticker := time.NewTicker(time.Minute)
	r.watchCh = r.cli.Watch(context.Background(), r.keyPrefix, clientv3.WithPrefix())

	for {
		select {
		case <-r.closeCh:
			return
		case res, ok := <-r.watchCh:
			if ok {
				r.update(res.Events)
			}
		case <-ticker.C:
			if err := r.sync(); err != nil {
				r.logger.Errorf("sync failed, err: %s", err)
			}
		}
	}
}
