package model

type Questionnaire struct {
	Title     string
	Questions []Question
	Result    map[string]string
}

type QuestionnairePayload struct {
	Questions []Question
	Title     string
}
