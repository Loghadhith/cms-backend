package post

import (
	"log"
	"net/http"

	"github.com/Loghadhith/cms/types"
	"github.com/Loghadhith/cms/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	post types.PostStore
}

func NewHandler(post types.PostStore) *Handler {
	return &Handler{post: post}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/postdata", h.handlePost).Methods("POST")
	router.HandleFunc("/fetchdata", h.getData).Methods("GET")

	// admin routes
	// router.HandleFunc("/users/{userID}", auth.WithJWTAuth(h.handleGetUser, h.store)).Methods(http.MethodGet)
}


func (h *Handler) handlePost(w http.ResponseWriter , r *http.Request){

  log.Println("post data")
  
  var fd types.PostPayload

  if err := utils.ParseJSON(r, &fd); err != nil {
    utils.WriteError(w, http.StatusNotFound,err)
    return
  }

  log.Println(fd.Data)
  log.Println(fd.Repo)
}


func (h *Handler) getData(w http.ResponseWriter, r * http.Request){
  log.Println("ok done")
}
