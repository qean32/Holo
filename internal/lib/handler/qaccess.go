package handler

import (
	"atomicgo.dev/cursor"

	errorlist "github.com/qean32/holo/internal/constants/error-list"
	"github.com/qean32/holo/internal/model"
	"github.com/qean32/holo/internal/terminal/questionnaire"
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
