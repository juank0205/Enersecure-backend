package models

import "time"

type Client struct {
	ClientID     uint       `db:"client_id" json:"client_id"`
	FirstName    string    `db:"first_name" json:"first_name"`
	LastName     string    `db:"last_name" json:"last_name"`
	GovernmentID uint     `db:"government_id" json:"government_id"`
	PhoneNumber  string    `db:"phone_number" json:"phone_number"`
	Password     string    `db:"password" json:"-"`
	Email        string    `db:"email" json:"email"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
