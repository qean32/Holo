package model

type Event struct {
	Key      string
	KeyWord  string
	Payload  string
	SubEvent *Event
	Password string
	Flags    []string
	DateTime string
}
