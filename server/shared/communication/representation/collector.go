package representation

import (
	"github.com/go-ozzo/ozzo-validation"
)

type CreateReportRequest struct {
	UniqueIdentifier string  `json:"uniqueIdentifier"`
	RunID            string  `json:"runId"`
	Latitude         float32 `json:"latitude"`
	Longitude        float32 `json:"longitude"`
	Height           float32 `json:"height"`
	ObjectsCaptured  uint64  `json:"objectsCaptured"`
}

func (r *CreateReportRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UniqueIdentifier, validation.Required),
		validation.Field(&r.RunID, validation.Required),
		validation.Field(&r.Latitude, validation.Required),
		validation.Field(&r.Longitude, validation.Required),
	)
}

type CreateReportResponse struct {
}
