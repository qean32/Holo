package lib

import (
	"main/callstack/event"
	"main/lib/handler/registry"
	"main/model"
)

func HandleEvent(e model.Event) error {
	if err := registry.Dispatch(e); err != nil {
		return fail(e, err)
	}

	if registry.Success(e.Key) {
		event.Success()
	}

	return logEvent(e)
}
