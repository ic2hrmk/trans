package main

import (
	"log"

	"trans/server/registry"
	"trans/server/shared/cmd"
)

//go:generate go run main.go --kind=route

func main() {
	//
	// Load startup flags
	//
	flags := cmd.LoadFlags()

	//
	// Load env.
	//
	if flags.EnvFile != "" {
		err := cmd.LoadEnvFile(flags.EnvFile)
		if err != nil {
			log.Fatal(err)
		}
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
	log.Fatal(service.Serve(flags.Address))
}
