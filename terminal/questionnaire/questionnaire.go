package questionnaire

import (
	"main/constants/sgr"
	"main/lib/utils"
	"main/model"
	"main/terminal"
	"main/terminal/field"
)

func Questionnaire(payload model.QuestionnairePayload) map[string]string {
	result := make(map[string]string, len(payload.Questions))

	terminal.Println(utils.GetCustomMessage(payload.Title+"\n", sgr.DIM))
	for index, question := range payload.Questions {
		isEnd := index == len(payload.Questions)-1
		if answer := askQuestion(question, isEnd); question.Callback(answer) {
			result[question.Key] = answer
		} else {
			newQuestion := question
			newQuestion.Message = "invalid response " + question.Message
			askQuestion(newQuestion, isEnd)
		}
	}
	terminal.DownAndStart()
	return result
}

func askQuestion(question model.Question, isEnd bool) string {
	return field.Field(model.FieldPayload{
		Prefix:    question.Message + " > ",
		PrefixSRG: []int{sgr.DIM},
		Transfer:  isEnd,
	})
}
