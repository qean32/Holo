package lib

import (
	"github.com/qean32/holo/internal/callstack/event"
	"github.com/qean32/holo/internal/lib/handler/registry"
	"github.com/qean32/holo/internal/model"
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
