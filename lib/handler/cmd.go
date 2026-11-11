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

func CreateBat(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.CreateFile(cmdPath(event.KeyWord), event.Payload); err != nil {
		return errorlist.Create(event.KeyWord, err)
	}
	return nil
}

func RunBat(event model.Event) error {
	if err := death.RunCmd(filepath.Join(constants.Root, cmdPath(event.KeyWord))); err != nil {
		return errorlist.Run(event.KeyWord, err)
	}
	return nil
}

func RemoveCmd(event model.Event) error {
	if err := death.RemoveFile(cmdPath(event.KeyWord)); err != nil {
		return errorlist.Remove(event.KeyWord, err)
	}
	return nil
}

func QuestionnaireCreateBat() error {
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
		Title: literal.Titles.EnterBat,
	})

	return CreateBat(model.Event{
		Key:      eventlist.BAT,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: death.CurrentTime(),
	})
}
