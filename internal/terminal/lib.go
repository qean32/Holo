package terminal

import (
	"fmt"
	"strings"

	"atomicgo.dev/cursor"
	"github.com/nathan-fiscaletti/consolesize-go"

	"github.com/qean32/holo/internal/constants"
)

func PrintASCIICenter(ascii string, separator string) {
	lines := strings.Split(ascii, "\n")
	if len(lines) < 2 {
		Print(ascii)
		return
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	offset := calcCenterCMD(maxLen - 2)
	if offset < 0 {
		offset = 0
	}

	repeat := strings.Repeat(separator, offset)
	length := len(lines) - 2

	params := make([]any, length)
	for i := range params {
		params[i] = repeat
	}

	Print(fmt.Sprintf(ascii, params...))
}

func DownAndStart() {
	Println()
	cursor.StartOfLine()
}

func ClearLines(count int) {
	for i := 0; i < count; i++ {
		cursor.ClearLine()
		cursor.Up(1)
	}
}

func calcCenterCMD(length int) int {
	cols, _ := consolesize.GetConsoleSize()
	res := cols/2 - length/2
	if res < 0 {
		return 0
	}
	return res
}

func RenderBanner() {
	PrintASCIICenter(constants.BinaryPROJECT_INIT, " ")
	PrintASCIICenter(constants.PROJECT_INIT, " ")
}
