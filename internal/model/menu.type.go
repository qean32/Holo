package model

type List struct {
	Title    string
	Position int
	Length   int
	Options  []Option
}

type ListPayload struct {
	Options []Option
	Title   string
}
