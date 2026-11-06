package list

import (
	"fmt"
	"strconv"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"main/callstack"
	"main/callstack/event"
	"main/constants/sgr"
	"main/lib/utils"
	"main/model"
	"main/terminal"
)

type state struct {
	position int
	options  []model.Option
}

func (s *state) length() int {
	return len(s.options)
}

func List(payload model.ListPayload) {
	terminal.DownAndStart()
	if len(payload.Options) == 0 {
		event.Response("length == zero! ave caesar! ")
		return
	}

	cursor.Hide()
	terminal.Println(utils.GetCustomMessage(" "+payload.Title+" ↓"+" (ctrl+c to back)", sgr.DIM))
	terminal.PrintTopLine()

	s := &state{options: payload.Options}
	s.render()
	s.jumpToStart()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if value, err := strconv.Atoi(key.String()); err == nil {
			s.jumpTo(value - 1)
		}
		switch key.Code {
		case keys.Enter:
			s.selectOption()
			return true, nil
		case keys.Down:
			s.moveDown()
		case keys.Up:
			s.moveUp()
		case keys.Escape, keys.CtrlC:
			cursor.Show()
			s.jumpToEnd()
			return true, nil
		}

		return false, nil
	})
	terminal.DownAndStart()
}

func (s *state) render() {
	for i, item := range s.options {
		isSelected := i == s.position

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

func (s *state) selectOption() {
	s.jumpToEnd()
	cursor.Show()
	callstack.Push(utils.GetShortEvent(s.options[s.position].Event))
}
