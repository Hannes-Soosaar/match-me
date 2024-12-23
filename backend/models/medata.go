package models

import "time"

type ProfileInformation struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Created   time.Time `json:"created_at"`
	City      string    `json:"user_city"`
	About     string    `json:"about_me"`
	Birthdate time.Time `json:"birthdate"`
}
