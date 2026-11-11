package lib

import (
	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	"main/lib/death"
	"main/model"
	"main/terminal"
)

func Init() error {
	if err := InitData(); err != nil {
		return err
	}

	terminal.RenderBanner()
	manual.Manual()
	return nil
}

func InitData() error {
	if err := constants.InitRoot(); err != nil {
		return err
	}
	return nil
}

func fail(e model.Event, err error) error {
	event.Response(err)
	if logErr := death.Logger(model.StackError{Event: e, Err: err}); logErr != nil {
		event.Response(logErr)
	}
	return err
}

func logg(e model.Event) error {
	if err := death.Logger(model.StackError{Event: e}); err != nil {
		event.Response(err)
		return err
	}
	return nil
}
