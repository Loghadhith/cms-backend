package api

import (
	"database/sql"
	"net/http"

	"github.com/Loghadhith/cms/services/post"
	"github.com/Loghadhith/cms/services/user"
	"github.com/gorilla/mux"
)


func Run(s *sql.DB , router *mux.Router) mux.Router {

  subrouter := router.PathPrefix("/api/v1").Subrouter()


	userStore := user.NewStore(s)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

  postStore := post.NewStore(s)
  postHandler := post.NewHandler(postStore)
  postHandler.RegisterRoutes(subrouter)

  router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))


  return *router
}
