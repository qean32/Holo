package handler

import (
	errorlist "main/constants/error-list"
	"main/model"
	"main/terminal/questionnaire"

	"atomicgo.dev/cursor"
)

func QAccess(event model.Event, warning string) error {
	cursor.Up(1)
	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message:  "confirm (yea/no)" + warning,
				Key:      "Response",
				Callback: func(string) bool { return true },
			},
		},
		Title: "",
	})

	if !questionnaire.IsAcceptableYea(q["Response"]) {
		return errorlist.ErrAccessDenied
	}

	return nil
}
