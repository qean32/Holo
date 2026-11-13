package event

import (
	"github.com/qean32/holo/internal/callstack"
	eventlist "github.com/qean32/holo/internal/constants/event-list"
	"github.com/qean32/holo/internal/lib/utils"
	"github.com/qean32/holo/internal/model"
)

func Response(message ...any) {
	callstack.Push(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: utils.ConcatMessage(message...),
	})
}

func Reset() {
	callstack.Push(model.Event{Key: eventlist.RESET})
}

func Success() {
	Response("/ᐠ-˕-マ Успех!")
}
