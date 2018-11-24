package neon_6m

import "github.com/ic2hrmk/go-event"

type neon6MR struct {
	devicePath string
}

func (r *neon6MR) Subscribe(coordinateEventStream *go_event.EventStream) {
	panic("implement me")
}

func (r *neon6MR) Run() error {
	panic("implement me")
}

func (r *neon6MR) GetCurrentPosition() (latitude, longitude float32, err error) {
	panic("implement me")
}
