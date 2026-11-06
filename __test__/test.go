package test

import (
	"fmt"
	"main/callstack"
	"main/constants"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/model"
	"path/filepath"
)

func Test() {
	callstack.PushMany(
		model.Event{Key: eventlist.TESTVIEW},
		model.Event{Key: eventlist.PUREREMOVEBAT, KeyWord: "test"},
		model.Event{Key: eventlist.RUNBAT, KeyWord: "test"},
		model.Event{Key: eventlist.BAT, KeyWord: "test", Payload: "start cmd /C echo 67"},
	)
}

func TestView() {
	if err := death.WriteFile(callstack.GetErrorsString(), filepath.Join("tmplog.txt")); err != nil {
		fmt.Println(err)
		return
	}

	defer callstack.PushMany(
		model.Event{Key: eventlist.PUREREMOVEBAT, KeyWord: "tmpbat"},
		model.Event{Key: eventlist.RUNBAT, KeyWord: "tmpbat"},
		model.Event{Key: eventlist.BAT, KeyWord: "tmpbat", Payload: "start " + filepath.Join(constants.Root, "tmplog.txt")},
	)
}
