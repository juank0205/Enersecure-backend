package app

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/juank0205/Enersecure-backend/internal/db"
	"github.com/juank0205/Enersecure-backend/internal/http/routes"
	"github.com/juank0205/Enersecure-backend/internal/http/server"
)

type EnvVariables struct {
	dbUser string
	dbPass string
	dbName string
	dbHost string

	serverPort uint64
}

type App struct {
	db        *db.DB
	variables EnvVariables
}

func loadEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var v EnvVariables
	v.dbUser = os.Getenv("DBUSER")
	v.dbPass = os.Getenv("DBPASS")
	v.dbName = os.Getenv("DBNAME")
	v.dbHost = os.Getenv("DBHOST")

	var serverPortStr = os.Getenv("SERVERPORT")
	v.serverPort, err = strconv.ParseUint(serverPortStr, 10, 0)
	if err != nil {
		log.Fatalf("invalid uint value for MY_UINT_VAR: %v", err)
	}
	return v
}

func RunApp() {
	var app App
	app.variables = loadEnvVariables()
	app.db = db.ConnectDB(
		app.variables.dbUser,
		app.variables.dbPass,
		app.variables.dbName,
		app.variables.dbHost,
	)

	mux := server.CreateMux()
	routes.SetupRoutes(app.db, mux)
	server.StartServer(app.db, mux, uint(app.variables.serverPort))
}
