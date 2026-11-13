package handler

import (
	"os"
	"path/filepath"

	"github.com/qean32/holo/internal/callstack/event"
	"github.com/qean32/holo/internal/callstack/manual"
	"github.com/qean32/holo/internal/constants"
	errorlist "github.com/qean32/holo/internal/constants/error-list"
	"github.com/qean32/holo/internal/constants/literal"
	"github.com/qean32/holo/internal/lib/death"
	"github.com/qean32/holo/internal/lib/global"
	"github.com/qean32/holo/internal/terminal"
)

const dirPerm os.FileMode = 0o755

func toggleResponse(enabled bool, onMsg, offMsg string) {
	if enabled {
		event.Response(onMsg)
		return
	}
	event.Response(offMsg)
}

func Reset() {
	death.ClearTerminal()
	terminal.RenderBanner()
	manual.Manual()
}

func Drop() error {
	if err := os.RemoveAll(constants.Root); err != nil {
		return errorlist.Remove(constants.Root, err)
	}
	if err := os.MkdirAll(constants.Root, dirPerm); err != nil {
		return errorlist.Create(constants.Root, err)
	}
	return nil
}

func ClearLog() error {
	if err := death.ClearFile(literal.Path.PathLog); err != nil {
		return errorlist.ClearLog(err)
	}
	return nil
}

func OpenLog() error {
	path := filepath.Join(constants.Root, literal.Path.PathLog)
	if err := death.RunScript(path); err != nil {
		return errorlist.Open(path, err)
	}
	return nil
}

func Inwork() {
	event.Response("В разработке")
}

func ToggleLogger() {
	toggleResponse(
		global.EventLogToggle(),
		"Логирование событий: включено",
		"Логирование событий: выключено",
	)
}

func ToggleScript() {
	toggleResponse(
		global.ScriptToggle(),
		"Работа со сценариями: включено",
		"Работа со сценариями: выключено",
	)
}

func InitSystem() error {
	if err := os.MkdirAll(filepath.Join(constants.Root, constants.Script), dirPerm); err != nil {
		return errorlist.Create(constants.Script, err)
	}
	return death.CreateFile(literal.Path.PathLog, "")
}
