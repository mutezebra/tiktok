package registry

import "github.com/sarulabs/di"

type Container struct {
	ctn  di.Container
	name string
}

func NewContainer(name string, build func(di.Container) (interface{}, error)) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	err = builder.Add([]di.Def{
		{
			Name:  name,
			Build: build,
		},
	}...)
	if err != nil {
		return nil, err
	}

	return &Container{
		ctn:  builder.Build(),
		name: name,
	}, nil
}

func (c *Container) Resolve() interface{} {
	return c.ctn.Get(c.name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}
