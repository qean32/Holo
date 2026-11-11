package registry

import (
	"main/callstack/manual"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/lib/handler"
	"main/lib/handler/menu"
	"main/model"
	test "main/test"
	"slices"
)

var commands = []model.Command{
	// сценарии
	{Name: eventlist.SCRIPT, Options: model.ParseOptions{KeyWord: true, Payload: true}, Fn: decorateScriptEvent(handler.CreateScript), Success: true},
	{Name: eventlist.RUNSCRIPT, Options: model.ParseOptions{KeyWord: true}, Fn: decorateScriptEvent(handler.RunScript), Success: true},
	{Name: eventlist.REMOVESCRIPT, Options: model.ParseOptions{SubEvent: eventlist.PUREREMOVESCRIPT, GetSubEvent: true}, NeedAccess: true, Fn: decorateScriptEvent(handler.RemoveScript)},
	{Name: eventlist.PUREREMOVESCRIPT, Fn: decorateScriptEvent(handler.RemoveScript), Success: true},
	{Name: eventlist.LISTSCRIPT, Fn: decorateScript(ignoreError(func() { list.ScriptList(eventlist.RUNSCRIPT) }))},
	{Name: eventlist.LISTSCRIPTREMOVE, Fn: decorateScript(ignoreError(func() { list.ScriptList(eventlist.REMOVESCRIPT) }))},
	{Name: eventlist.SCRIPTTOGGLE, Fn: ignoreError(handler.ToggleScript)},
	// core
	{Name: eventlist.MENU, Fn: ignoreError(list.MainMenu)},
	{Name: eventlist.DROP, Fn: ignoreEvent(handler.Drop), Success: true, NeedAccess: true},
	{Name: eventlist.TEST, Fn: ignoreError(test.Test), Success: true},
	{Name: eventlist.STOP, Fn: ignoreError(death.Exit)},
	{Name: eventlist.HELP, Fn: ignoreError(handler.Help)},
	// logs
	{Name: eventlist.OPENLOG, Fn: ignoreEvent(handler.ClearLog)},
	{Name: eventlist.CLEARLOG, Fn: ignoreEvent(handler.ClearLog), Success: true},
	{Name: eventlist.LOGGERTOGGLE, Fn: ignoreError(handler.ToggleLogger)},
	// Служебные команды и события меню
	{Name: eventlist.INIT, Fn: ignoreEvent(handler.InitSystem), Success: true, NeedAccess: true},
	{Name: eventlist.INWORK, Fn: ignoreError(handler.Inwork)},
	{Name: eventlist.RESPONSE, Fn: handler.Response},
	{Name: eventlist.MANUAL, Fn: ignoreError(manual.Manual)},
	{Name: eventlist.TESTVIEW, Fn: ignoreError(test.TestView)},
	{Name: eventlist.RESET, Fn: ignoreError(handler.Reset)},
	// Опросники
	{Name: eventlist.QCREATESCRIPT, Fn: ignoreEvent(handler.QuestionnaireCreateScript)},
}

var registry = buildRegistry(commands)

func Dispatch(e model.Event) (model.Action, error) {
	command, ok := registry[e.Key]

	if command.Options.GetSubEvent {
		e.SubEvent = &model.Event{Key: command.Options.SubEvent, KeyWord: e.KeyWord}
	}

	if !ok {
		return nil, errorlist.ErrUnknownCommand
	}
	if command.NeedAccess && !slices.Contains(e.Flags, "-y") {
		if err := handler.QAccess(e, ""); err != nil {
			return nil, err
		}
	}
	return nil, command.Fn(e)
}
