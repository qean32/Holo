package registry

import (
	"slices"

	"github.com/qean32/holo/internal/callstack/manual"
	errorlist "github.com/qean32/holo/internal/constants/error-list"
	eventlist "github.com/qean32/holo/internal/constants/event-list"
	"github.com/qean32/holo/internal/lib/death"
	"github.com/qean32/holo/internal/lib/handler"
	"github.com/qean32/holo/internal/lib/handler/menu"
	"github.com/qean32/holo/internal/model"
	test "github.com/qean32/holo/script"
)

var commands = []model.Command{
	// сценарии
	{Name: eventlist.SCRIPT, Options: model.ParseOptions{KeyWord: true, Payload: true}, Fn: decorateScriptEvent(handler.CreateScript), Success: true},
	{Name: eventlist.RUNSCRIPT, Options: model.ParseOptions{KeyWord: true}, Fn: decorateScriptEvent(handler.RunScript), Success: true},
	{Name: eventlist.REMOVESCRIPT, Options: model.ParseOptions{SubEvent: eventlist.PUREREMOVESCRIPT, GetSubEvent: true}, NeedAccess: true, Fn: decorateScriptEvent(handler.RemoveScript)},
	{Name: eventlist.PUREREMOVESCRIPT, Fn: decorateScriptEvent(handler.RemoveScript), Success: true},
	{Name: eventlist.LISTSCRIPT, Fn: decorateScript(ignoreError(func() { menu.ScriptList(eventlist.RUNSCRIPT) }))},
	{Name: eventlist.LISTSCRIPTREMOVE, Fn: decorateScript(ignoreError(func() { menu.ScriptList(eventlist.REMOVESCRIPT) }))},
	{Name: eventlist.SCRIPTTOGGLE, Fn: ignoreError(handler.ToggleScript)},
	// core
	{Name: eventlist.MENU, Fn: ignoreError(menu.MainMenu)},
	{Name: eventlist.DROP, Fn: ignoreEvent(handler.Drop), Success: true, NeedAccess: true},
	{Name: eventlist.STOP, Fn: ignoreError(death.Exit)},
	{Name: eventlist.HELP, Fn: ignoreError(handler.Help)},
	// логи
	{Name: eventlist.OPENLOG, Fn: ignoreEvent(handler.ClearLog)},
	{Name: eventlist.CLEARLOG, Fn: ignoreEvent(handler.ClearLog), Success: true},
	{Name: eventlist.LOGGERTOGGLE, Fn: ignoreError(handler.ToggleLogger)},
	// Служебные команды и события меню
	{Name: eventlist.INIT, Fn: ignoreEvent(handler.InitSystem), Success: true, NeedAccess: true},
	{Name: eventlist.INWORK, Fn: ignoreError(handler.Inwork)},
	{Name: eventlist.RESPONSE, Fn: handler.Response},
	{Name: eventlist.MANUAL, Fn: ignoreError(manual.Manual)},
	{Name: eventlist.RESET, Fn: ignoreError(handler.Reset)},
	// Опросники
	{Name: eventlist.QCREATESCRIPT, Fn: ignoreEvent(handler.QuestionnaireCreateScript)},
	// test
	{Name: eventlist.TESTVIEW, Fn: ignoreError(test.TestView)},
	{Name: eventlist.TEST, Fn: ignoreError(test.Test)},
}

var registry = buildRegistry(commands)

func Dispatch(e model.Event) error {
	command, ok := registry[e.Key]

	if command.Options.GetSubEvent {
		e.SubEvent = &model.Event{Key: command.Options.SubEvent, KeyWord: e.KeyWord}
	}

	if !ok {
		return errorlist.ErrUnknownCommand
	}
	if command.NeedAccess && !slices.Contains(e.Flags, "-y") {
		if err := handler.QAccess(e, ""); err != nil {
			return err
		}
	}
	return command.Fn(e)
}
