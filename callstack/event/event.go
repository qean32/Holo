package event

import (
	"main/callstack"
	eventlist "main/constants/event-list"
	"main/lib/utils"
	"main/model"
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
