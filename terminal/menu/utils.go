package menu

import (
	"atomicgo.dev/cursor"

	"main/constants"
)

func (state *stateStruct) jumpTo(newPos int) {
	if newPos < 0 || newPos >= state.length() {
		return
	}
	diff := newPos - state.position
	if diff > 0 {
		cursor.Down(diff)
	} else if diff < 0 {
		cursor.Up(-diff)
	}
	state.position = newPos
	state.rerender()
}

func (state *stateStruct) moveUp() {
	if state.decrementPosition() {
		state.rerender()
	}
}

func (state *stateStruct) moveDown() {
	if state.incrementPosition() {
		state.rerender()
	}
}

func (state *stateStruct) incrementPosition() bool {
	if state.position < state.length()-1 {
		state.position++
		cursor.Down(1)
		return true
	}
	if state.length() > 1 {
		state.position = 0
		state.jumpToStart()
		cursor.Down(1)
		return true
	}
	return false
}

func (state *stateStruct) jumpToEnd() {
	cursor.Down(state.length() - state.position)
}

func (state *stateStruct) jumpToStart() {
	cursor.Up(state.length() - state.position)
}

func getStartSymbol(isSelected bool) string {
	if isSelected {
		return constants.CHAR_SELECTED_ITEM
	}
	return constants.CHAR_UN_SELECTED_ITEM
}

func (state *stateStruct) decrementPosition() bool {
	if state.position > 0 {
		state.position--
		cursor.Up(1)
		return true
	}
	if state.length() > 1 {
		diff := state.length() - 1 - state.position
		if diff > 0 {
			cursor.Down(diff)
		}
		state.position = state.length() - 1
		return true
	}
	return false
}
