package models

import "time"

type ClientAddress struct {
	AddressID int       `db:"client_address_id" json:"client_address_id"`
	ClientID  int       `db:"client_id" json:"client_id"`
	CityID    uint      `db:"city_id" json:"city_id"`
	Address   string    `db:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
