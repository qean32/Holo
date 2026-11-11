package model

type Menu struct {
	Title    string
	Position int
	Length   int
	Options  []Option
}

type ListPayload struct {
	Options []Option
	Title   string
}
