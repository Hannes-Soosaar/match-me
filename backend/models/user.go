package models

import "time"

type User struct {
	ID               string    `json:"id"`
	Uuid             string    `json:"uuid"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"password_hash"`
	CreatedAt        time.Time `json:"created_at"`
	UserCity         string    `json:"user_city"`
	RegisterLocation string    `json:"register_location"`
	BrowsLocation    string    `json:"brows_location"`
}
