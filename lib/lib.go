package lib

import (
	"slices"

	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/lib/handler"
	"main/model"
	"main/terminal"
)

var successEvents = []string{
	eventlist.SCRIPT,
	eventlist.PUREREMOVESCRIPT,
	eventlist.TEST,
	eventlist.CLEARLOG,
	eventlist.CRYPTO,
	eventlist.DROP,
	eventlist.INIT,
	eventlist.RUNSCRIPT,
}

func Init() {
	if err := constants.InitRoot(); err != nil {
		terminal.Println("Ошибка инициализации")
		return
	}
	terminal.RenderBanner()
	manual.Manual()
}

func HandleEvent(e model.Event) error {
	if err := handler.Dispatch(e); err != nil {
		event.Response(err)
		if logErr := death.Logger(model.StackError{Event: e, Err: err}); logErr != nil {
			event.Response(logErr)
		}
		return err
	}

	if slices.Contains(successEvents, e.Key) {
		event.Success()
	}

	if e.Key == eventlist.RESET {
		return nil
	}

	if err := death.Logger(model.StackError{Event: e}); err != nil {
		event.Response(err)
		return err
	}
	return nil
}
