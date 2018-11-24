package event_storage

import (
	"time"
	"trans/client/app/persistence/history"
	"trans/client/app/persistence/history/model"
	"trans/client/app/persistence/history/repository"
)

type EventStorage struct {
	videoEventRepository  repository.VideoEventRepository
	gpsEventsRepository   repository.GPSEventRepository
	errorEventsRepository repository.ErrorEventRepository
}

func NewEventStorage(
	videoEventRepository repository.VideoEventRepository,
	gpsEventsRepository repository.GPSEventRepository,
	errorEventsRepository repository.ErrorEventRepository,
) *EventStorage {
	return &EventStorage{
		videoEventRepository:  videoEventRepository,
		gpsEventsRepository:   gpsEventsRepository,
		errorEventsRepository: errorEventsRepository,
	}
}

func (rcv *EventStorage) AddVideoEvent(e *history.VideoLogEvent) error {
	_, err := rcv.videoEventRepository.Create(&model.VideoEventRecord{
		CreatedAt:      time.Now().Unix(),
		RunID:          e.RunID,
		ObjectsCounter: e.ObjectsCounter,
	})

	if err != nil {
		return err
	}

	return nil
}

func (rcv *EventStorage) AddGPSEvent(e *history.GPSLogEvent) error {
	_, err := rcv.gpsEventsRepository.Create(&model.GPSEventRecord{
		CreatedAt: time.Now().Unix(),
		RunID:     e.RunID,
		Longitude: e.Longitude,
		Latitude:  e.Latitude,
		Height:    e.Height,
	})

	if err != nil {
		return err
	}

	return nil
}

func (rcv *EventStorage) AddErrorEvent(e *history.ErrorLogEvent) error {
	_, err := rcv.errorEventsRepository.Create(&model.ErrorEventRecord{
		CreatedAt: time.Now().Unix(),
		RunID:     e.RunID,
		Message:   e.Message,
	})

	if err != nil {
		return err
	}

	return nil
}

func (rcv *EventStorage) DeleteRunLogs(runID string) error {

	if err := rcv.videoEventRepository.DeleteByRunID(runID); err != nil {
		return err
	}

	if err := rcv.gpsEventsRepository.DeleteByRunID(runID); err != nil {
		return err
	}

	if err := rcv.errorEventsRepository.DeleteByRunID(runID); err != nil {
		return err
	}

	return nil
}

func (rcv *EventStorage) DeleteLogsOlderThen(timestamp time.Time) error {

	if err := rcv.videoEventRepository.DeleteOlderThen(timestamp); err != nil {
		return err
	}

	if err := rcv.gpsEventsRepository.DeleteOlderThen(timestamp); err != nil {
		return err
	}

	if err := rcv.errorEventsRepository.DeleteOlderThen(timestamp); err != nil {
		return err
	}

	return nil
}
