package manual

import (
	"strings"

	"github.com/qean32/holo/internal/callstack"
	"github.com/qean32/holo/internal/callstack/event"
	"github.com/qean32/holo/internal/constants"
	errorlist "github.com/qean32/holo/internal/constants/error-list"
	"github.com/qean32/holo/internal/lib/handler/parse"
	"github.com/qean32/holo/internal/model"
	"github.com/qean32/holo/internal/terminal/field"
)

func Manual() {
	command := field.Field(model.FieldPayload{Prefix: constants.FIELDPREFIX})
	args := strings.Fields(command)
	if len(args) == 0 {
		event.Response(errorlist.ErrEmptyCommand)
		return
	}

	e, err := parse.Dispatch(args[0], args)
	if err != nil {
		event.Response(err)
		return
	}
	callstack.Push(e)
}
