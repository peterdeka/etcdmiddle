package etcdmiddle

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
)

type EtcdCli struct {
	Client *client.Client
}

var instanceEtcdCli *EtcdCli = nil

func Connect(servers []string) (*client.Client, error) {
	if instanceEtcdCli == nil {
		cfg := client.Config{
			Endpoints: []string{"http://127.0.0.1:2379"},
			Transport: client.DefaultTransport,
			// set timeout per request to fail fast when the target endpoint is unavailable
			HeaderTimeoutPerRequest: time.Second,
		}
		c, err := client.New(cfg)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		instanceEtcdCli = &EtcdCli{Client: &c}
	}
	return instanceEtcdCli.Client, nil
}
