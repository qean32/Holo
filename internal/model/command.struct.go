package model

type Command struct {
	Name       string
	Options    ParseOptions
	Fn         Handler
	Success    bool
	NeedAccess bool
}
