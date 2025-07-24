package server

import (
	"fmt"
	"net/http"

	"github.com/juank0205/Enersecure-backend/internal/db"
)

func CreateMux() *http.ServeMux {
	return http.NewServeMux()
}

func StartServer(db *db.DB, mux *http.ServeMux, port uint) {
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("Server listening on port %d\n", port)
	s.ListenAndServe()
}
