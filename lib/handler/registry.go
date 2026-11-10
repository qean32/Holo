package handler

import (
	test "main/__test__"
	"main/callstack/manual"
	"main/lib/death"
	"main/lib/global"
	"main/model"

	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
)

var registry = map[string]model.Handler{
	// клиентские
	eventlist.INIT: ignoreEvent(init_),

	eventlist.HELP:   ignoreError(help),
	eventlist.MENU:   ignoreError(runMenu),
	eventlist.RESET:  ignoreError(reset),
	eventlist.DROP:   ignoreEvent(drop),
	eventlist.STOP:   ignoreError(death.Exit),
	eventlist.MANUAL: ignoreError(manual.Manual),

	eventlist.LOGGERTOGGLE: ignoreError(toggleLogger),
	eventlist.SCRIPTTOGGLE: ignoreError(toggleScript),

	eventlist.CLEARLOG: ignoreEvent(clearLog),
	eventlist.OPENLOG:  ignoreEvent(openLog),

	eventlist.CRYPTO:   ignoreEvent(encrypt),
	eventlist.DECRYPTO: ignoreEvent(decrypt),

	eventlist.GENERATEKEY:    ignoreEvent(generateKey),
	eventlist.GENERATEMASTER: ignoreEvent(generateMasterKey),

	// технические
	eventlist.INWORK:     ignoreError(inwork),
	eventlist.TEST:       ignoreError(test.Test),
	eventlist.TESTVIEW:   ignoreError(test.TestView),
	eventlist.RESPONSE:   _response,
	eventlist.MENUCRYPTO: ignoreError(inwork),
	// сценарии
	eventlist.PUREREMOVESCRIPT: decorateScriptEvent(removeCmd),
	eventlist.SCRIPT:           decorateScriptEvent(createScript),
	eventlist.MANUCREATESCRIPT: decorateScript(questionnaireCreateBat),

	eventlist.RUNSCRIPT:    decorateScriptEvent(runScript),
	eventlist.REMOVESCRIPT: decorateScriptEvent(qAccess),

	eventlist.LISTSCRIPT:       decorateScript(runBatList),
	eventlist.LISTSCRIPTREMOVE: decorateScript(removeBatList),
}

func ignoreError(fn func()) model.Handler {
	return func(_ model.Event) error {
		fn()
		return nil
	}
}

func ignoreEvent(fn func() error) model.Handler {
	return func(_ model.Event) error {
		return fn()
	}
}

func decorateScript(fn func() error) model.Handler {
	return func(_ model.Event) error {
		if !global.ScriptEnabled() {
			return errorlist.ErrScriptOff
		}
		return fn()
	}
}

func decorateScriptEvent(fn func(e model.Event) error) model.Handler {
	return func(e model.Event) error {
		if !global.ScriptEnabled() {
			return errorlist.ErrScriptOff
		}
		return fn(e)
	}
}

func Dispatch(e model.Event) error {
	fn, ok := registry[e.Key]
	if !ok {
		return errorlist.ErrUnknownCommand
	}
	return fn(e)
}
