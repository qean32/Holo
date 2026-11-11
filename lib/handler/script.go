package handler

import (
	"path/filepath"
	"strings"

	"main/constants"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/death"
	"main/model"
	"main/terminal"
	"main/terminal/questionnaire"
)

func Response(event model.Event) error {
	terminal.PrintResponse(event.Payload)
	return nil
}

func CreateScript(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.CreateFile(scriptPath(event.KeyWord), event.Payload); err != nil {
		return errorlist.Create(event.KeyWord, err)
	}
	return nil
}

func RunScript(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.RunCmd(filepath.Join(constants.Root, scriptPath(event.KeyWord))); err != nil {
		return errorlist.Run(event.KeyWord, err)
	}
	return nil
}

func RemoveScript(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.RemoveFile(scriptPath(event.KeyWord)); err != nil {
		return errorlist.Remove(event.KeyWord, err)
	}
	return nil
}

func QuestionnaireCreateScript() error {
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

	return CreateScript(model.Event{
		Key:      eventlist.SCRIPT,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: death.CurrentTime(),
	})
}
