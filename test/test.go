package test

import (
	"fmt"
	"path/filepath"

	"main/callstack"
	"main/constants"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/model"
)

func Test() {
	callstack.PushMany(
		model.Event{Key: eventlist.TESTVIEW},
		model.Event{Key: eventlist.PUREREMOVESCRIPT, KeyWord: "test"},
		model.Event{Key: eventlist.RUNSCRIPT, KeyWord: "test"},
		model.Event{Key: eventlist.SCRIPT, KeyWord: "test", Payload: "start cmd /C echo 67"},
	)
}

func TestView() {
	if err := death.WriteFile(callstack.GetErrorsString(), filepath.Join("tmplog.txt")); err != nil {
		fmt.Println(err)
		return
	}

	defer callstack.PushMany(
		model.Event{Key: eventlist.PUREREMOVESCRIPT, KeyWord: "tmpbat"},
		model.Event{Key: eventlist.RUNSCRIPT, KeyWord: "tmpbat"},
		model.Event{Key: eventlist.SCRIPT, KeyWord: "tmpbat", Payload: "start " + filepath.Join(constants.Root, "tmplog.txt")},
	)
}
