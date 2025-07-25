package models

import "time"

type ClientTurbineRent struct {
	RentID     int       `db:"rent_id" json:"rent_id"`
	ClientID   int       `db:"client_id" json:"client_id"`
	AddressID  int       `db:"address_id" json:"address_id"`
	TurbineID  int       `db:"turbine_id" json:"turbine_id"`
	Amount     int       `db:"amount" json:"amount"`
	StartDate  time.Time `db:"start_date" json:"start_date"`
	IsActive   bool      `db:"is_active" json:"is_active"`
	EndDate    time.Time `db:"end_date" json:"end_date"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
