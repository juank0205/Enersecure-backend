package models

import "time"

type Invoice struct {
	InvoiceID int       `db:"invoice_id" json:"invoice_id"`
	ClientID  int       `db:"client_id" json:"client_id"`
	Date      time.Time `db:"date" json:"date"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
