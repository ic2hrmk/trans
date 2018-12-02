package repository

import (
	"errors"
	"trans/server/app/vehicle/persistence/model"
)

var ErrClassNotFound = errors.New("ERR_CLASS_NOT_FOUND")

type ClassRepository interface {
	CreateClass(class *model.Class) error
	GetClassByID(classID string) (*model.Class, error)
}
