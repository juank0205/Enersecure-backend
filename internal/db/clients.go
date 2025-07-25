package db

import (
	"database/sql"
	"fmt"

	"github.com/juank0205/Enersecure-backend/internal/models"
)

// Uses the email and password to check for matches in the database
// Returns the id of the matching row on a successful login and an error on no match
func (db *DB) GetEmailAndPassword(email string, password string) (int64, error) {
	var clientId int64
	row := db.sql.QueryRow("SELECT client_id FROM clients where email = ? AND password = ?", email, password)
	if err := row.Scan(&clientId); err != nil {
		return int64(clientId), fmt.Errorf("GetEmailAndPassword: %v", err)
	}
	return clientId, nil
}

// Checks if the email provided is already associated to a client register
// Returns true if the email is available and false if there is already a client with that email
func (db *DB) CheckEmailAvailable(email string) (bool, error) {
	var clientEmail string
	row := db.sql.QueryRow("SELECT email FROM clients WHERE email = ?", email)
	err := row.Scan(&clientEmail)

	if err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		return false, fmt.Errorf("CheckEmailAvailable: %v", err)
	}
return false, nil }


// Creates a client in the database and returns its client_id
func (db *DB) RegisterClient(c *models.Client) (int64, error) {
	query := `INSERT INTO clients (
		first_name, last_name, government_id, phone_number, password, email
	) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := db.sql.Exec(query,
		c.FirstName,
		c.LastName,
		c.GovernmentID,
		c.PhoneNumber,
		c.Password,
		c.Email,
	)
	if err != nil {
		return 0, fmt.Errorf("registerClient: %v", err)
	}
	return result.LastInsertId()
}
