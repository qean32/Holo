package qcallback

import (
	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"github.com/qean32/holo/internal/callstack/event"
	"github.com/qean32/holo/internal/terminal"
)

func QCallback(callback func()) {
	defer cursor.Show()
	terminal.Print("press Enter to continue")
	cursor.Hide()
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.Enter:
			callback()
			return true, nil
		case keys.Escape, keys.CtrlC:
			event.Reset()
			return true, nil
		}

		return false, nil
	})
}
