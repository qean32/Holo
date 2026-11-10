package handler

import (
	"path/filepath"
	"strings"

	"main/constants"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/array"
	"main/lib/death"
	"main/model"
	"main/terminal"
	"main/terminal/list"
	"main/terminal/questionnaire"
)

func _response(event model.Event) error {
	terminal.PrintResponse(event.Payload)
	return nil
}

func createScript(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.CreateFile(cmdPath(event.KeyWord), event.Payload); err != nil {
		return errorlist.Create(event.KeyWord, err)
	}
	return nil
}

func runScript(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.RunCmd(filepath.Join(constants.Root, cmdPath(event.KeyWord))); err != nil {
		return errorlist.Run(event.KeyWord, err)
	}
	return nil
}

func removeCmd(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.RemoveFile(cmdPath(event.KeyWord)); err != nil {
		return errorlist.Remove(event.KeyWord, err)
	}
	return nil
}

func listBatFiles(title string, makeEvent func(name string) model.Event) error {
	dir := filepath.Join(constants.Root, constants.Script)
	commands, err := death.ListFilesByExt(dir, literal.Extension.Script)
	if err != nil {
		return errorlist.Read(dir, err)
	}

	list.List(model.ListPayload{
		Options: array.Map(commands, func(v string) model.Option {
			return model.Option{
				Message: strings.TrimSuffix(v, literal.Extension.Script),
				Event:   makeEvent(v),
			}
		}),
		Title: title,
	})
	return nil
}

func runBatList() error {
	return listBatFiles(literal.Titles.ListScript, func(value string) model.Event {
		return model.Event{
			Key:     eventlist.RUNSCRIPT,
			KeyWord: trimExtension(value),
		}
	})
}

func removeBatList() error {
	return listBatFiles(literal.Titles.ListScript, func(value string) model.Event {
		trimmed := trimExtension(value)
		return model.Event{
			Key:     eventlist.REMOVESCRIPT,
			KeyWord: trimmed,
			SubEvent: &model.Event{
				Key:     eventlist.PUREREMOVESCRIPT,
				KeyWord: trimmed,
			},
		}
	})
}

func questionnaireCreateBat() error {
	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message: "Ключ",
				Key:     "KeyWord",
				Callback: func(res string) bool {
					return validateName(res) == nil
				},
			},
			{
				Message: "Команда",
				Key:     "Payload",
				Callback: func(res string) bool {
					return strings.TrimSpace(res) != ""
				},
			},
		},
		Title: literal.Titles.EnterScript,
	})

	return createScript(model.Event{
		Key:      eventlist.SCRIPT,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: death.CurrentTime(),
	})
}
