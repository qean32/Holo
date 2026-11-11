package handler

import (
	"os"
	"path/filepath"

	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	errorlist "main/constants/error-list"
	"main/constants/literal"
	"main/lib/death"
	"main/lib/global"
	"main/terminal"
)

const dirPerm os.FileMode = 0o755

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
	if err := death.RunCmd(path); err != nil {
		return errorlist.Open(path, err)
	}
	return nil
}

func Inwork() {
	event.Response("В разработке")
}

func ToggleLogger() {
	if global.EventLogToggle() {
		event.Response("Логирование событий: включено")
		return
	}
	event.Response("Логирование событий: выключено")
}

func ToggleScript() {
	if global.ScriptToggle() {
		event.Response("Работа со сценариями: включено")
		return
	}
	event.Response("Работа со сценариями: выключено")
}

func InitSystem() error {
	if err := os.MkdirAll(filepath.Join(constants.Root, constants.Script), dirPerm); err != nil {
		return errorlist.Create(constants.Script, err)
	}
	return death.CreateFile(literal.Path.PathLog, "")
}
