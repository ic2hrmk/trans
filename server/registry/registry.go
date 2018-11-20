package registry

import (
	"trans/server/app/collector"
	"trans/server/app/prover"
	"trans/server/app/route"
	"trans/server/app/vehicle"
)

var c Container

func init() {
	c = Container{}

	c.add(collector.ServiceName, collector.FactoryMethod)
	c.add(prover.ServiceName, prover.FactoryMethod)
	c.add(route.ServiceName, route.FactoryMethod)
	c.add(vehicle.ServiceName, vehicle.FactoryMethod)
}

func Get(serviceName string) (FactoryMethod, error) {
	return c.get(serviceName)
}
