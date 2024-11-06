package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/norrico31/it210-auth-service-backend/config"
	"github.com/norrico31/it210-auth-service-backend/service/user"
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

func (s *APIServer) enforceGatewayOrigin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Construct allowed host based on the config
		allowedHost := fmt.Sprintf("%s:%s", s.config.PublicHost, s.config.GatewayPort)
		fmt.Print(allowedHost)
		// Check if the request is coming from the specified gateway origin
		if r.Host != allowedHost && !strings.HasPrefix(r.Referer(), fmt.Sprintf("http://%s", allowedHost)) {
			fmt.Println("IT'S NOT ACCESSING HERE ")
			http.Error(w, "NOT FOUND", http.StatusNotFound)
			return
		}

		// Allow request to proceed if origin is correct
		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	// addr := fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port)
	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		// TODO: ADJUST CONDITION BASE ON PROD/DEV URL OR PATH
	// 		fmt.Println("HSOT:: ", r.Host)
	// 		fmt.Println("PORT FROM ENV:: ", addr)
	// 		fmt.Println("s.Address:: ", s.addr)
	// 		if addr == r.Host {
	// 			http.Error(w, "Forbidden", http.StatusForbidden)
	// 			return
	// 		}

	// 		next.ServeHTTP(w, r)
	// 	})
	// })
	router.Use(s.enforceGatewayOrigin)

	subrouterv1 := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	user.RegisterRoutes(subrouterv1, userHandler)

	log.Println("Auth Service: Running on port ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
