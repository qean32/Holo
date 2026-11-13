package model

type Action interface {
	Run() ([]Event, error)
}

type ActionFunc func() ([]Event, error)

func (f ActionFunc) Run() ([]Event, error) {
	return f()
}
