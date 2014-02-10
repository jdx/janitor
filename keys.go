package janitor

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

type Getter interface {
	Get(key string, sort, recursive bool) (*etcd.Response, error)
	Watch(prefix string, waitIndex uint64, recursive bool, receiver chan *etcd.Response, stop chan bool) (*etcd.Response, error)
}

func GetKeys(client Getter) string {
	resp, err := client.Get("authorized_keys", false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current keys: %s %s", resp.Node.Key, resp.Node.Value)
	return resp.Node.Value
}

//watchChan := make(chan *etcd.Response)
//go client.Watch("/authorized_keys", 0, false, watchChan, nil)
//log.Println("Waiting for an update...")
//r := <-watchChan
//log.Println("Got updated creds:", r.Node)
