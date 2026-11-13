package callstack

import (
	"fmt"
	"strings"
	"sync"

	"github.com/qean32/holo/internal/model"
)

var (
	errorsMu sync.Mutex
	Errors   []model.StackError
)

func PushError(event model.Event, err error) {
	if err == nil {
		return
	}
	errorsMu.Lock()
	defer errorsMu.Unlock()
	Errors = append(Errors, model.StackError{Event: event, Err: err})
}

func GetErrorsString() string {
	errorsMu.Lock()
	defer errorsMu.Unlock()

	if len(Errors) == 0 {
		return "no error"
	}

	var sb strings.Builder
	for i, stackErr := range Errors {
		sb.WriteString(StackErrorToString(stackErr, i))
	}
	return sb.String()
}

func StackErrorToString(stackErr model.StackError, i int) string {
	return fmt.Sprintf("[%d] Event: %+v, Error: %v\n", i+1, stackErr.Event, stackErr.Err)
}
