package parse

import (
	"strings"

	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/lib/array/filter"
	"main/lib/death"
	"main/model"
)

const (
	prefixPassword = "-pass:"
	prefixPayload  = "-p:"
)

var registry = map[string]model.ParseOptions{
	eventlist.CRYPTO:   {KeyWord: true, Payload: true, Password: true},
	eventlist.DECRYPTO: {KeyWord: true, Password: true},

	eventlist.SCRIPT:       {KeyWord: true, Payload: true},
	eventlist.RUNSCRIPT:    {KeyWord: true},
	eventlist.REMOVESCRIPT: {GetSubEvent: true, SubEvent: eventlist.REMOVESCRIPT, KeyWord: true},
}

func Dispatch(key string, args []string) (model.Event, error) {
	return parse(args, registry[key])
}

func parse(args []string, opts model.ParseOptions) (model.Event, error) {
	if len(args) == 0 {
		return model.Event{}, errorlist.ErrEmptyCommand
	}

	event := model.Event{
		Key:      args[0],
		DateTime: death.CurrentTime(),
		Flags:    filter.FilterIsFlag(args),
	}

	if opts.KeyWord {
		if len(args) < 2 {
			return event, errorlist.ErrTooFewArgs
		}
		event.KeyWord = args[1]
	}

	if opts.Password {
		pass, err := getItem(args, prefixPassword)
		if err != nil {
			return event, err
		}
		event.Password = pass
	}

	if opts.Payload {
		payload, err := getPayload(args, prefixPayload)
		if err != nil {
			return event, err
		}
		event.Payload = payload
	}

	if opts.GetSubEvent {
		event.SubEvent = &model.Event{
			KeyWord: args[1],
			Key:     opts.SubEvent,
		}
	}

	return event, nil
}

// getItem ищет в аргументах первое вхождение с заданным префиксом
// и возвращает значение после префикса.
func getItem(args []string, prefix string) (string, error) {
	for _, a := range args {
		if v, ok := strings.CutPrefix(a, prefix); ok {
			return v, nil
		}
	}
	return "", errorlist.ErrNotExist
}

// getPayload извлекает полезную нагрузку: всё после первого вхождения префикса,
// включая остальные аргументы.
func getPayload(args []string, prefix string) (string, error) {
	for i, a := range args {
		if v, ok := strings.CutPrefix(a, prefix); ok {
			rest := make([]string, 0, len(args)-i)
			rest = append(rest, v)
			rest = append(rest, args[i+1:]...)
			return strings.TrimSpace(strings.Join(rest, " ")), nil
		}
	}
	return "", errorlist.ErrNoPayload
}
