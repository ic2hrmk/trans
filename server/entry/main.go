package main

import (
	"log"

	"trans/server/registry"
	"trans/server/shared/init"
)

func main() {
	//
	// Load startup flags
	//
	flags := init.LoadFlags()

	//
	// Load env.
	//
	if flags.EnvFile != "" {
		init.LoadEnvFile(flags.EnvFile)
	}

	//
	// Select service
	//
	serviceFactory, err := registry.Get(flags.Kind)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Create service
	//
	service, err := serviceFactory()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Run till the death comes
	//
	log.Fatal(service.Serve())
}
