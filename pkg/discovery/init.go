package discovery

import (
	"time"

	etcd "go.etcd.io/etcd/client/v3"

	"github.com/Mutezebra/tiktok/config"
)

func newClient() (*etcd.Client, error) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{config.Conf.Etcd.Endpoint},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
