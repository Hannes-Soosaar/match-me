package models

type MultipleChoiceQuestion struct {
	ID int `json:"id"`
	Questions []string `json:"questions"`
	Answer int 	`json:"answer"`
}



