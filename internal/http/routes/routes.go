package routes

import (
	"net/http"

	"github.com/juank0205/Enersecure-backend/internal/db"
	"github.com/juank0205/Enersecure-backend/internal/http/handlers"
	"github.com/juank0205/Enersecure-backend/internal/http/middlewares"
)

func SetupRoutes(db *db.DB, rootMux *http.ServeMux) {
	h := handlers.NewHandler(db)
	publicStack := middlewares.CreateStack(
		middlewares.Logging,
		middlewares.Recover,
	)

	privateStack := middlewares.CreateStack(
		middlewares.Logging,
		middlewares.Recover,
		middlewares.AuthMiddleware,
	)

	publicMux := http.NewServeMux()
	publicMux.HandleFunc("POST /login", h.HandleLogin)
	publicMux.HandleFunc("POST /register", h.HandleRegister)

	privateMux := http.NewServeMux()
	// privateMux.HandleFunc("GET /test", h.HandleTest)
	//
	// // Clients
	// privateMux.HandleFunc("GET /clients/search", h.HandleGetClientsByNameMatch)
	// privateMux.HandleFunc("POST /clients", h.HandleCreateClient)
	//
	// // Movings
	// privateMux.HandleFunc("GET /movings", h.HandleGetAllMovings)
	// privateMux.HandleFunc("POST /movings/create", h.HandleCreateMoving)
	// privateMux.HandleFunc("GET /movings/get", h.HandleGetMovingById)
	//
	// //Moving_workers
	// privateMux.HandleFunc("POST /moving_workers/assign", h.HandleAssignMovingWorkers)
	//
	// //Workers
	// privateMux.HandleFunc("GET /workers", h.HandleGetWorkers)
	//
	// //Storages
	// privateMux.HandleFunc("POST /storages/create", h.HandleCreateStorage)
	// privateMux.HandleFunc("GET /storages", h.HandleGetStorages)
	//
	rootMux.Handle("/api/v1/public/", http.StripPrefix("/api/public", publicStack(publicMux)))
	rootMux.Handle("/api/v2/private/", http.StripPrefix("/api/private", privateStack(privateMux)))
}
