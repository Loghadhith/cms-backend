package api

import (
	"database/sql"
	"net/http"

	"github.com/Loghadhith/cms/services/post"
	"github.com/Loghadhith/cms/services/user"
	"github.com/gorilla/mux"
)

// type APIServer struct {
// 	addr string
// 	db   *sql.DB
// }

// func NewAPIServer(addr string, db *sql.DB) *APIServer {
// 	return &APIServer{
// 		addr: addr,
// 		db:   db,
// 	}
// }

func Run(s *sql.DB , router *mux.Router) mux.Router {
	// subrouter := router.PathPrefix("/api/v1").Subrouter()

  subrouter := router.PathPrefix("/api/v1").Subrouter()


	userStore := user.NewStore(s)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

  postStore := post.NewStore(s)
  postHandler := post.NewHandler(postStore)
  postHandler.RegisterRoutes(subrouter)

  router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

  // log.Println("")
  // log.Println(s.addr)

	// log.Println("Listening on:1", s.addr)
 //  log.Println(userHandler)

  // srv := http.Server{
  //   Handler: router,
  //   Addr: s.addr,
  // }

	// srv.ListenAndServe()
  return *router
}
