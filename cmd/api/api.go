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

// TODO: STILL NOT WORKING IN DOCKER
func (s *APIServer) enforceGatewayOrigin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Construct allowed host based on the config
		allowedHost := fmt.Sprintf("%s:%s", s.config.PublicHost, s.config.GatewayPort)
		fmt.Printf("rHost: %s", r.Host)
		fmt.Printf("gatewayPort: %s", s.config.GatewayPort)
		if r.Host == allowedHost {
			// Allow requests that come from the gateway (127.0.0.1:8080)
			next.ServeHTTP(w, r)
			return
		}

		// If the request is directly to the auth service (127.0.0.1:8081), return NOT FOUND
		if r.Host == fmt.Sprintf("127.0.0.1:%s", s.config.GatewayPort) {
			http.Error(w, "NOT FOUND", http.StatusNotFound)
			return
		}

		// Optional: Check the referer header as additional verification
		if !strings.HasPrefix(r.Referer(), fmt.Sprintf("http://%s", allowedHost)) {
			http.Error(w, "NOT FOUND", http.StatusNotFound)
			return
		}

		// Allow request to proceed if it's from the correct gateway
		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.Use(s.enforceGatewayOrigin)

	subrouterv1 := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	user.RegisterRoutes(subrouterv1, userHandler)

	log.Println("Auth Service: Running on port ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
