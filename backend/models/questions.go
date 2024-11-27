package models

// This is a general bio, this will not be scored

type Questions struct {
	ID        int    `json:"id"`
	Questions string `json:"question"`
	Answer    string `json:"answer"`
}
