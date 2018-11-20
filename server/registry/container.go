package registry

import (
	"fmt"

	"trans/server/app"
)

type Container map[string]FactoryMethod

type FactoryMethod func() (app.MicroService, error)

func (c *Container) add(serviceName string, fabric FactoryMethod) {
	(*c)[serviceName] = fabric
}

func (c *Container) get(serviceName string) (FactoryMethod, error) {
	entry, ok := (*c)[serviceName]
	if !ok {
		return nil, fmt.Errorf("service [%s] isn't registered", serviceName)
	}

	return entry, nil
}
