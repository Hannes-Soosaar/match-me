package models

// Send a list of interests to the Front end so they can be selected.
// The interest name as pre-made

type Interests struct {
	ID           int    `json:"id"`
	CategoryID   int    `json:"categoryID"`
	InterestName string `json:"interestName"`
}
