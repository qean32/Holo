package list

import (
	"atomicgo.dev/cursor"

	"github.com/qean32/holo/internal/constants"
	"github.com/qean32/holo/internal/terminal"
)

func (s *state) jumpToEnd() {
	cursor.Down(s.length() - s.position)
}

func (s *state) jumpToStart() {
	cursor.Up(s.length() - s.position)
}

func getStartSymbol(isSelected bool) string {
	if isSelected {
		return constants.CHAR_SELECTED_ITEM
	}
	return constants.CHAR_UN_SELECTED_ITEM
}

func (s *state) jumpTo(newPos int) {
	if newPos < 0 || newPos >= s.length() {
		return
	}
	diff := newPos - s.position
	if diff > 0 {
		cursor.Down(diff)
	} else if diff < 0 {
		cursor.Up(-diff)
	}
	s.position = newPos
	s.rerender()
}

func (s *state) moveUp() {
	if s.decrementPosition() {
		s.rerender()
	}
}

func (s *state) moveDown() {
	if s.incrementPosition() {
		s.rerender()
	}
}

func (s *state) incrementPosition() bool {
	if s.position < s.length()-1 {
		s.position++
		cursor.Down(1)
		return true
	}
	if s.length() > 1 {
		s.position = 0
		s.jumpToStart()
		cursor.Down(1)
		return true
	}
	return false
}

func (s *state) decrementPosition() bool {
	if s.position > 0 {
		s.position--
		cursor.Up(1)
		return true
	}
	if s.length() > 1 {
		diff := s.length() - 1 - s.position
		if diff > 0 {
			cursor.Down(diff)
		}
		s.position = s.length() - 1
		return true
	}
	return false
}

func (s *state) rerender() {
	s.jumpToEnd()
	terminal.ClearLines(s.length())
	s.render()
	s.jumpToStart()
}
