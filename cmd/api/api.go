package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/norrico31/it210-auth-service-backend/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
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

	subrouterv1 := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	user.RegisterRoutes(subrouterv1, userHandler)

	log.Println("Auth Service: Running on port ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
