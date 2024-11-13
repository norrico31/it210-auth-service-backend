package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/norrico31/it210-auth-service-backend/config"
	"github.com/norrico31/it210-auth-service-backend/services/user"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	config config.Config
}

func NewApiServer(addr string, db *sql.DB, cfg config.Config) *APIServer {
	return &APIServer{
		addr:   addr,
		db:     db,
		config: cfg,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouterv1 := router.PathPrefix("/api/v1/auth").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	user.RegisterRoutes(subrouterv1, userHandler)

	log.Println("Auth Service: Running on port ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
