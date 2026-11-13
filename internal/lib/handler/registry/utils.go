package registry

import (
	errorlist "github.com/qean32/holo/internal/constants/error-list"
	"github.com/qean32/holo/internal/lib/global"
	"github.com/qean32/holo/internal/model"
)

func Success(name string) bool {
	command, ok := registry[name]
	return ok && command.Success
}

func buildRegistry(list []model.Command) map[string]model.Command {
	result := make(map[string]model.Command, len(list))
	for _, command := range list {
		if command.Name == "" {
			panic("invalid command: empty name")
		}
		if _, exists := result[command.Name]; exists {
			panic("duplicate command: " + command.Name)
		}
		result[command.Name] = command
	}
	return result
}

func ignoreError(fn func()) model.Handler {
	return func(_ model.Event) error {
		fn()
		return nil
	}
}

func ignoreEvent(fn func() error) model.Handler {
	return func(_ model.Event) error {
		return fn()
	}
}

func decorateScript(fn model.Handler) model.Handler {
	return func(e model.Event) error {
		if !global.ScriptEnabled() {
			return errorlist.ErrScriptOff
		}
		return fn(e)
	}
}

func decorateScriptEvent(fn func(e model.Event) error) model.Handler {
	return decorateScript(fn)
}
