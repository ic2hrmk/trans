package route_service

import (
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type RouteService struct {
	routeRepository repository.RouteRepository
	runRepository   repository.RunRepository
}

func NewRouteService(
	routeRepository repository.RouteRepository,
	runRepository repository.RunRepository,
) *RouteService {
	return &RouteService{
		routeRepository: routeRepository,
		runRepository:   runRepository,
	}
}

func (rcv *RouteService) StartRun(routeID string) error {
	//
	// Finish current run if there is one
	//
	if err := rcv.StopCurrentRun(); err != nil {
		return err
	}

	//
	// Create new run
	//
	if _, err := rcv.runRepository.Create(&model.Run{
		RouteID: routeID,
		Status:  model.RunStatusActive,
	}); err != nil {
		return err
	}

	return nil
}

func (rcv *RouteService) StopCurrentRun() error {
	//
	// Finish all active runs
	//
	activeRuns, err := rcv.runRepository.FindByStatus(model.RunStatusActive)
	if err != nil {
		return err
	}

	for i := range activeRuns {
		activeRuns[i].Status = model.RunStatusEnded

		if _, err := rcv.runRepository.Update(activeRuns[i]); err != nil {
			return err
		}
	}

	return nil
}

func (rcv *RouteService) IsCurrentRunActive() (bool, error) {
	activeRuns, err := rcv.runRepository.FindByStatus(model.RunStatusActive)
	if err != nil {
		return false, err
	}

	return len(activeRuns) != 0, nil
}
