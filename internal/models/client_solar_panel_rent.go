package models

import "time"

type ClientSolarPanelRent struct {
	RentID        int       `db:"rent_id" json:"rent_id"`
	ClientID      int       `db:"client_id" json:"client_id"`
	AddressID     int       `db:"address_id" json:"address_id"`
	SolarPanelID  int       `db:"solar_pannel_id" json:"solar_pannel_id"`
	Amount        int       `db:"amount" json:"amount"`
	StartDate     time.Time `db:"start_date" json:"start_date"`
	IsActive      bool      `db:"is_active" json:"is_active"`
	EndDate       time.Time `db:"end_date" json:"end_date"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}
