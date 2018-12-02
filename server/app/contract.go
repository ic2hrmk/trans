package app

type MicroService interface {
	Serve(address string) error
}
