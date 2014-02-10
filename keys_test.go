package janitor_test

import (
	"github.com/coreos/go-etcd/etcd"
	. "github.com/dickeyxxx/janitor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeClient struct{}

func (c fakeClient) Get(key string, sort, recursive bool) (*etcd.Response, error) {
	resp := new(etcd.Response)
	resp.Node = new(etcd.Node)
	resp.Node.Key = "foo"
	resp.Node.Value = "bar"
	return resp, nil
}

func (c fakeClient) Watch(_ string, _ uint64, _ bool, _ chan *etcd.Response, _ chan bool) (*etcd.Response, error) {
	resp := new(etcd.Response)
	return resp, nil
}

var _ = Describe("Keys", func() {
	It("Works!", func() {
		client := new(fakeClient)
		Ω(GetKeys(client)).To(Equal("bar"))
	})
})
