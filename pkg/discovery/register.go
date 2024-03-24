package discovery

import (
	"context"
	"time"

	"github.com/pkg/errors"
	etcd "go.etcd.io/etcd/client/v3"

	"github.com/Mutezebra/tiktok/pkg/log"
)

type Registry struct {
	Addr   string
	TTL    int64
	Key    string
	Prefix string

	leaseID       etcd.LeaseID
	client        *etcd.Client
	keepAliveChan <-chan *etcd.LeaseKeepAliveResponse
}

func NewRegistry(Addr string, key string, TTL int64, prefix ...string) (*Registry, error) {
	if len(prefix) > 1 {
		return nil, errors.New("the size of prefix must be 0 or 1")
	}
	pre := ""
	if prefix != nil {
		pre = prefix[0]
		if pre[len(pre)-1] != '/' {
			pre = pre + "/"
		}
	}
	if TTL == 0 {
		TTL = 15
	}
	client, err := newClient()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create the etcd client")
	}
	return &Registry{
		Addr:          Addr,
		TTL:           TTL,
		Key:           key,
		Prefix:        pre,
		leaseID:       0,
		keepAliveChan: nil,
		client:        client,
	}, nil
}

// Register registers the service to etcd and starts the keepAlive goroutine.
func (r *Registry) Register(ctx context.Context) error {
	grantCtx, cn := context.WithTimeout(ctx, 1*time.Second)
	defer cn()
	lease, err := r.client.Grant(grantCtx, r.TTL)
	if err != nil {
		return errors.WithMessage(err, "grant lease failed")
	}
	r.leaseID = lease.ID

	putCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	key := ""
	if r.Prefix != "" {
		key = r.Prefix + r.Key
	}
	_, err = r.client.Put(putCtx, key, r.Addr, etcd.WithLease(r.leaseID))

	if err != nil {
		_ = r.client.Close()
		return errors.WithMessage(err, "etcd client put server failed")
	}
	go r.keepAlive(ctx)
	return nil
}

func (r *Registry) MustRegister(ctx context.Context) {
	if err := r.Register(ctx); err != nil {
		panic(err)
	}
}

// keepAlive maintains the service registration in etcd by sending keep-alive messages.
func (r *Registry) keepAlive(ctx context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	r.keepAliveChan, _ = r.client.KeepAlive(childCtx, r.leaseID)
	defer func() { _ = r.client.Close() }()
	for {
		select {
		case <-r.keepAliveChan:
			break
		case <-time.After(time.Duration(r.TTL) * time.Second):
			log.LogrusObj.Errorf("the service whose key is %s and address is %s loses its heartbeat in etcd", r.Key, r.Addr)
			return
		}
	}
}

func (r *Registry) Close() {
	_ = r.client.Close()
}
