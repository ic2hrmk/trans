package generators

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"

	"trans/server/app/route"
	"trans/server/app/route/persistence/model"
	"trans/server/app/route/persistence/repository/mongo"
	"trans/server/shared/utils"
)

func generateRoute() *model.Route {
	const routeNameTemplate = "Маршрут %s №%d"

	var vehicleTypes = []string{
		"тролейбусу", "трамаваю", "автобусу",
	}

	routeName := fmt.Sprintf(routeNameTemplate,
		vehicleTypes[rand.Int31n(int32(len(vehicleTypes)))], rand.Int31n(100))

	return &model.Route{
		RouteID: uuid.New().String(),
		Name:    routeName,
		Length:  5 + 10*rand.Float32(),
		StartPoint: model.RoutePoint{
			Latitude:  50 + rand.Float32(),
			Longitude: 30 + rand.Float32(),
		},
		EndPoint: model.RoutePoint{
			Latitude:  50 + rand.Float32(),
			Longitude: 30 + rand.Float32(),
		},
		Schedule: []*model.ScheduleSection{{
			From:     6*time.Hour + 30*time.Minute,
			To:       21*time.Hour + 45*time.Minute,
			Duration: time.Minute * time.Duration(5+rand.Int31n(10)),
		}},
	}
}

func CreateRoute(mongoURL string) error {
	routeDB, err := utils.InitMongoDBConnection(mongoURL, route.ServiceName)
	if err != nil {
		return err
	}

	return mongo.NewRouteRepository(routeDB).CreateRoute(generateRoute())
}
