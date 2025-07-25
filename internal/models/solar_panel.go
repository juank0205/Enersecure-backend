package models

import "time"

type SolarPanel struct {
	PanelID         int       `db:"pannel_id" json:"pannel_id"`
	Model           string    `db:"model" json:"model"`
	MaxCapacity     int       `db:"max_capacity" json:"max_capacity"`
	MaintainanceRate int      `db:"maintainance_rate" json:"maintainance_rate"`
	Efficiency      float32   `db:"efficiency" json:"efficiency"`
	RentPrice       float32   `db:"rent_price" json:"rent_price"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
