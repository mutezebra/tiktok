package inject

import (
	"github.com/Mutezebra/tiktok/app/domain/model"
	"github.com/Mutezebra/tiktok/pkg/discovery"
)

type Registry struct {
	Key    string
	Prefix string
	TTL    int64
	Addr   string
}

func NewRegistry(registry *Registry) (model.Registry, error) {
	return discovery.NewRegistry(registry.Addr, registry.Key, registry.TTL, registry.Prefix)
}

type Resolver struct {
}

func NewResolver() (model.Resolver, error) {
	return discovery.NewResolver()
}
