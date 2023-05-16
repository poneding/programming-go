package main

import (
	"fmt"
	"log"
	"os"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

var etcdEndpoints []string

func main() {

	etcdcli, err := etcd.New(etcd.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: 5 * time.Second,
	})
	failOnError(err, "Failed to new etcd client")

	_, err = etcdcli.Put(etcdcli.Ctx(), "foo", "bar1")
	failOnError(err, "Failed to put foo")

	_, err = etcdcli.Put(etcdcli.Ctx(), "foo", "bar2")
	failOnError(err, "Failed to put foo")

	gr, err := etcdcli.Get(etcdcli.Ctx(), "foo")
	failOnError(err, "Failed to get foo")

	for _, kv := range gr.Kvs {
		fmt.Println(string(kv.Key), string(kv.Value))
	}

	_, err = etcdcli.Delete(etcdcli.Ctx(), "foo")
	failOnError(err, "Failed to delete foo")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	etcdHost := os.Getenv("ETCD_HOST")
	if etcdHost == "" {
		log.Fatalln("ETCD_HOST environment variable is required")
	}
	etcdPort := os.Getenv("ETCD_PORT")
	if etcdPort == "" {
		etcdPort = "2379"
	}
	etcdEndpoints = []string{
		fmt.Sprintf("http://%s:%s", etcdHost, etcdPort),
	}
}
