package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/juank0205/Enersecure-backend/internal/db"
)

func CreateMux() *http.ServeMux {
	return http.NewServeMux()
}

func StartServer(db *db.DB, mux *http.ServeMux, port uint) {
	fmt.Printf("Server listening on port %d\n", port)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "cert/cert.pem", "cert/key.pem", mux)
	if err != nil {
		log.Fatalf("Failed to start HTTPS server: %v\n", err)
	}
}
