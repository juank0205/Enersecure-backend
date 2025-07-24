package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DB struct {
	sql *sql.DB
}

func ConnectDB(user string, password string, name string, host string) *DB {
	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = password
	cfg.Net = "tcp"
	cfg.Addr = host
	cfg.DBName = name

	//get a database handle.
	sqldb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Connection failed")
		log.Fatal(err)
	}

	pingErr := sqldb.Ping()
	if pingErr != nil {
		fmt.Println("Ping Failed")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return  &DB{sql: sqldb}
}
