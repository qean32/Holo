package menu

import (
	"fmt"
	"strconv"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"main/callstack"
	"main/callstack/event"
	"main/constants"
	"main/constants/sgr"
	"main/lib/utils"
	"main/model"
	"main/terminal"
)

type stateStruct struct {
	position int
	options  []model.Option
}

func (state *stateStruct) length() int {
	return len(state.options)
}

func Menu(payload model.ListPayload) (model.Event, bool) {
	terminal.DownAndStart()
	if len(payload.Options) == 0 {
		event.Response("length == zero! ave caesar! ")
		return model.Event{}, false
	}
	var selected model.Event
	chosen := false

	cursor.Hide()
	terminal.Println(utils.GetCustomMessage(payload.Title+" ↓"+" (ctrl+c back)", constants.StyleDim...))
	terminal.PrintTopLine()

	state := &stateStruct{options: payload.Options}
	state.render()
	state.jumpToStart()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if value, err := strconv.Atoi(key.String()); err == nil {
			state.jumpTo(value - 1)
		}
		switch key.Code {
		case keys.Enter:
			state.selectOption()
			selected = utils.GetShortEvent(state.options[state.position].Event)
			chosen = true
			return true, nil
		case keys.Down:
			state.moveDown()
		case keys.Up:
			state.moveUp()
		case keys.Escape, keys.CtrlC:
			cursor.Show()
			state.jumpToEnd()
			return true, nil
		}

		return false, nil
	})
	terminal.DownAndStart()
	return selected, chosen
}

func (state *stateStruct) render() {
	for i, item := range state.options {
		isSelected := i == state.position

		color := sgr.DIM
		if isSelected {
			color = sgr.GREEN
		}
		startChar := utils.GetCustomMessage(getStartSymbol(isSelected), color)
		label := utils.GetCustomMessage(fmt.Sprintf("%d. %s", i+1, item.Message), sgr.DIM)

		cursor.ClearLine()
		if isSelected {
			terminal.Print(utils.GetCustomMessage("│", sgr.DIM), startChar+label)
		} else {
			terminal.Print(startChar + label)
		}
		terminal.DownAndStart()
	}
}

func (state *stateStruct) rerender() {
	state.jumpToEnd()
	terminal.ClearLines(state.length())
	state.render()
	state.jumpToStart()
}

func (state *stateStruct) selectOption() {
	state.jumpToEnd()
	cursor.Show()
	callstack.Push(utils.GetShortEvent(state.options[state.position].Event))
}
