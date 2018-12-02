package prover

import (
	"trans/server/app"
	"trans/server/app/prover/service"
)

const ServiceName = "prover"

func FactoryMethod() (app.MicroService, error) {
	return service.NewProverService(), nil
}
