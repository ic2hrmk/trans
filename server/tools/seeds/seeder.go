package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"
	"trans/server/tools/seeds/generators"
)

//go:generate go run seeder.go --seed=route --env=../../.env

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var seedType string
	var envFile string

	flag.StringVar(&seedType, "seed", "", "type of seed")
	flag.StringVar(&envFile, "env", "", "env file")
	flag.Parse()

	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			log.Fatal(err)
		}
	}

	var err error

	switch seedType {
	case "route":
		err = generators.CreateRoute(getMongoURL())
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("OK!")
}

func getMongoURL() string {
	return os.Getenv("MONGO_URL")
}
