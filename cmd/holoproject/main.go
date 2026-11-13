package main

import (
	"fmt"
	"runtime"

	"github.com/qean32/holo/internal/callstack"
	"github.com/qean32/holo/internal/callstack/event"
	eventlist "github.com/qean32/holo/internal/constants/event-list"
	"github.com/qean32/holo/internal/lib"
	"github.com/qean32/holo/internal/lib/global"
	"github.com/qean32/holo/internal/terminal"
	"github.com/qean32/holo/internal/terminal/qcallback"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("Эта программа поддерживается только на Windows.")
		return
	}

	lib.Init()
	go runLoop()

	select {}
}

func runLoop() {
	for {
		for {
			ev, ok := callstack.Pop()
			if !ok {
				break
			}
			if global.EventLogEnabled() && ev.Key != eventlist.RESPONSE {
				terminal.PrintTechInfo(ev)
			}
			if err := lib.HandleEvent(ev); err != nil {
				callstack.PushError(ev, err)
			}
		}
		qcallback.QCallback(func() { event.Reset() })
		<-callstack.Changed()
	}
}
