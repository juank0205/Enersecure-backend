package models

import "time"

type Turbine struct {
	TurbineID        int       `db:"turbine_id" json:"turbine_id"`
	Model            string    `db:"model" json:"model"`
	Height           float32   `db:"height" json:"height"`
	MaintainanceRate int       `db:"maintainance_rate" json:"maintainance_rate"`
	Efficiency       float32   `db:"efficiency" json:"efficiency"`
	RentPrice        float32   `db:"rent_price" json:"rent_price"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
