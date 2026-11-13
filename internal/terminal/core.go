package terminal

import (
	"fmt"
	"sync"

	"github.com/qean32/holo/internal/lib/utils"
)

var mu sync.Mutex

func Print(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Print(" " + utils.ConcatMessage(messages...))
}

func Println(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(" " + utils.ConcatMessage(messages...))
}

func PrintResponse(messages ...any) {
	Println()
	Println("[ИНФО] ", utils.ConcatMessage(messages...))
	Println()
}
func PrintTopLine() {
	Println("")
}
