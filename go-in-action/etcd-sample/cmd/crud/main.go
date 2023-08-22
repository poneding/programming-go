package main

import (
	"fmt"
	"log"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

func main() {
	fmt.Println("etcd example")
	etcdcli, err := etcd.New(etcd.Config{
		Endpoints: []string{
			"http://localhost:12379",
		},
		DialTimeout: 5 * time.Second,
	})
	failOnError(err, "Failed to new etcd client")
	defer etcdcli.Close()

	// 1、写入键值 --------------------------------------------------------
	_, err = etcdcli.Put(etcdcli.Ctx(), "foo", "bar1")
	failOnError(err, "Failed to put foo")
	_, err = etcdcli.Put(etcdcli.Ctx(), "foo", "bar2")
	failOnError(err, "Failed to put foo")
	// 前缀键值
	etcdcli.Put(etcdcli.Ctx(), "foo_prefix/key1", "bar1")
	etcdcli.Put(etcdcli.Ctx(), "foo_prefix/key2", "bar2")

	// 2、读取键值 --------------------------------------------------------
	gr, err := etcdcli.Get(etcdcli.Ctx(), "foo")
	failOnError(err, "Failed to get foo")
	for _, kv := range gr.Kvs {
		fmt.Println(string(kv.Key), string(kv.Value))
	}
	// 前缀键值
	get, err := etcdcli.Get(etcdcli.Ctx(), "foo_prefix", etcd.WithPrefix())
	failOnError(err, "Failed to get foo_prefix")
	for _, kv := range get.Kvs {
		fmt.Println(string(kv.Key), string(kv.Value))
	}

	// 3、删除键值 --------------------------------------------------------
	_, err = etcdcli.Delete(etcdcli.Ctx(), "foo")
	failOnError(err, "Failed to delete foo")
	// 前缀键值
	_, err = etcdcli.Delete(etcdcli.Ctx(), "foo_prefix", etcd.WithPrefix())
	failOnError(err, "Failed to delete foo_prefix")

	// 4、Watch 监听
	go func() {
		// etcdcli.Watch(etcdcli.Ctx(), "foo")                           // 指定键值
		watchRespC := etcdcli.Watch(etcdcli.Ctx(), "foo_watch_prefix", etcd.WithPrefix()) // 指定前缀
		for resp := range watchRespC {
			// 事件类型只有 PUT 和 DELETE
			for _, ev := range resp.Events {
				fmt.Printf("%s \t%q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()
	// 监听测试操作
	time.Sleep(1)
	etcdcli.Put(etcdcli.Ctx(), "foo_watch_prefix/key1", "bar1")
	etcdcli.Put(etcdcli.Ctx(), "foo_watch_prefix/key1", "bar2")
	etcdcli.Delete(etcdcli.Ctx(), "foo_watch_prefix", etcd.WithPrefix())

	// 5、KeepAlive --------------------------------------------------------
	// 租约，过期时间为5秒
	lease, err := etcdcli.Grant(etcdcli.Ctx(), 5)
	failOnError(err, "Failed to grant lease")
	// 包含租约的键值
	_, err = etcdcli.Put(etcdcli.Ctx(), "foo_lease", "bar", etcd.WithLease(lease.ID))
	failOnError(err, "Failed to put foo_lease")
	// 保持租约（单次）
	// 租约必须在5秒内续约，否则租约将过期，过期的租约将不再续约
	// tick := time.NewTicker(3 * time.Second)
	// for {
	// 	select {
	// 	case <-tick.C:
	// 		_, err = etcdcli.KeepAliveOnce(etcdcli.Ctx(), lease.ID)
	// 		failOnError(err, "Failed to keep alive once")
	// 	}
	// }

	// 保持租约（永久）
	// 租约必须在5秒内续约，否则租约将过期，过期的租约将不再续约

	// 不规范的 KeepAlive 写法
	// _, err = etcdcli.KeepAlive(etcdcli.Ctx(), lease.ID)
	// 没有接收 KeepAlive 返回的 Channel ，会导致 Channel 充满，后续的响应无法继续写入，而被丢弃
	// for {
	// 	time.Sleep(1 * time.Second)
	// }

	// 规范的 KeepAlive 写法
	keepAliveRespC, err := etcdcli.KeepAlive(etcdcli.Ctx(), lease.ID)
	failOnError(err, "Failed to keep alive")
	// 必须消费 keepAliveRespC
	for resp := range keepAliveRespC {
		if resp == nil {
			fmt.Println("lease expired")
			break
		}
		fmt.Println(resp.TTL)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
