package lib

import (
	"main/callstack"
	"main/callstack/event"
	"main/lib/handler/registry"
	"main/model"
)

func HandleEvent(e model.Event) error {
	action, err := registry.Dispatch(e)
	if err != nil {
		return fail(e, err)
	}

	if action != nil {
		events, err := action.Run()
		if err != nil {
			return fail(e, err)
		}
		for _, next := range events {
			callstack.Push(next)
		}
	}

	if registry.Success(e.Key) {
		event.Success()
	}

	return logg(e)
}
