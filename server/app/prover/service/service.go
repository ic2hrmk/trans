package service

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"trans/server/app"
	"trans/server/app/prover/dto"
)

type ProverService struct {
	WebContainer *restful.Container
}

func NewProverService() app.MicroService {
	service := &ProverService{
		WebContainer: restful.NewContainer(),
	}

	service.init()

	return service
}

func (rcv *ProverService) init() {
	ws := &restful.WebService{}

	ws.Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/experiment").
		To(rcv.createExperiment).
		Operation("createExperiment").
		Reads(dto.CreateExperimentRequest{}).
		Writes(dto.CreateExperimentResponse{}).
		Returns(200, http.StatusText(http.StatusOK), dto.CreateExperimentResponse{}).
		Returns(500, http.StatusText(http.StatusInternalServerError), dto.Error{}))

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept",},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		CookiesAllowed: false,
		Container:      rcv.WebContainer,
	}

	rcv.WebContainer.Filter(cors.Filter)
	rcv.WebContainer.Filter(rcv.WebContainer.OPTIONSFilter)
}

func (rcv *ProverService) Serve(address string) error {
	return http.ListenAndServe(address, rcv.WebContainer)
}
