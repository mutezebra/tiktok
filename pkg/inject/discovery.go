package inject

import (
	"github.com/Mutezebra/tiktok/app/domain/model"
	"github.com/Mutezebra/tiktok/pkg/discovery"
	"github.com/Mutezebra/tiktok/pkg/log"
)

type Registry struct {
	Key    string
	Prefix string
	TTL    int64
	Addr   string
}

func NewRegistry(registry *Registry) model.Registry {
	reg, err := discovery.NewRegistry(registry.Addr, registry.Key, registry.TTL, registry.Prefix)
	if err != nil {
		log.LogrusObj.Panic(err)
		return nil
	}
	return reg
}

type Resolver struct {
}

func NewResolver() (model.Resolver, error) {
	return discovery.NewResolver()
}
