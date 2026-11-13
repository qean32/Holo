package model

type Question struct {
	Message  string
	Key      string
	Callback func(res string) bool
}
