package lib

import (
	"github.com/qean32/holo/internal/callstack/event"
	"github.com/qean32/holo/internal/callstack/manual"
	"github.com/qean32/holo/internal/constants"
	"github.com/qean32/holo/internal/lib/death"
	"github.com/qean32/holo/internal/model"
	"github.com/qean32/holo/internal/terminal"
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

func logEvent(e model.Event) error {
	if err := death.Logger(model.StackError{Event: e}); err != nil {
		event.Response(err)
		return err
	}
	return nil
}
