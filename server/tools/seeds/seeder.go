package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"trans/server/app/route"
	"trans/server/app/vehicle"
	"trans/server/shared/utils"

	routeModel "trans/server/app/route/persistence/model"
	routePersistence "trans/server/app/route/persistence/repository/mongo"
	vehicleModel "trans/server/app/vehicle/persistence/model"
	vehiclePersistence "trans/server/app/vehicle/persistence/repository/mongo"
)

//go:generate go run seeder.go --seed=troll --env=../../.env

const (
	tram  = "tram"
	troll = "troll"
	bus   = "bus"
)

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

	routeEntity := generateRoute(seedType)
	classEntity := generateClass(seedType)
	vehicleEntity := generateVehicle()

	vehicleEntity.RouteID = routeEntity.RouteID
	vehicleEntity.ClassID = classEntity.ClassID

	routeDB, err := utils.InitMongoDBConnection(getMongoURL(), route.ServiceName)
	if err != nil {
		log.Fatal(err)
	}

	vehicleDB, err := utils.InitMongoDBConnection(getMongoURL(), vehicle.ServiceName)
	if err != nil {
		log.Fatal(err)
	}

	if err := routePersistence.NewRouteRepository(routeDB).CreateRoute(routeEntity); err != nil {
		log.Fatal(err)
	}

	if err := vehiclePersistence.NewClassRepository(vehicleDB).CreateClass(classEntity); err != nil {
		log.Fatal(err)
	}

	if err := vehiclePersistence.NewVehicleRepository(vehicleDB).CreateVehicle(vehicleEntity); err != nil {
		log.Fatal(err)
	}

	log.Println("OK!")
}

func getMongoURL() string {
	return os.Getenv("MONGO_URL")
}

func generateClass(vehicleType string) *vehicleModel.Class {
	return &vehicleModel.Class{
		ClassID: uuid.New().String(),
		Name:    "ЛАЗ E183",
		Type:    vehicleType,
		Vendor:  "ЛАЗ",
		Model:   "E183",
		Seats:   36,
		Stands:  120,
	}
}

func generateVehicle() *vehicleModel.Vehicle {
	return &vehicleModel.Vehicle{
		VehicleID:         uuid.New().String(),
		UniqueIdentifier:  string(strings.ToUpper(strings.Replace(uuid.New().String(), "-", "", -1))[10]),
		RegistrationPlate: fmt.Sprintf("%d", 1000+rand.Int31n(9000)),
		VIN:               strings.ToUpper(strings.Replace(uuid.New().String(), "-", "", -1)),
	}
}

func generateRoute(routeType string) *routeModel.Route {
	const routeNameTemplate = "Маршрут %s №%d"

	routeNameType := map[string]string{
		tram:  "трамаваю",
		troll: "тролейбусу",
		bus:   "автобусу",
	}[routeType]

	routeName := fmt.Sprintf(routeNameTemplate, routeNameType, rand.Int31n(100))

	return &routeModel.Route{
		RouteID: uuid.New().String(),
		Name:    routeName,
		Length:  5 + 10*rand.Float32(),
		StartPoint: routeModel.RoutePoint{
			Latitude:  50 + rand.Float32(),
			Longitude: 30 + rand.Float32(),
		},
		EndPoint: routeModel.RoutePoint{
			Latitude:  50 + rand.Float32(),
			Longitude: 30 + rand.Float32(),
		},
		Schedule: []*routeModel.ScheduleSection{{
			From:     6*time.Hour + 30*time.Minute,
			To:       21*time.Hour + 45*time.Minute,
			Duration: time.Minute * time.Duration(5+rand.Int31n(10)),
		}},
	}
}
