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
	router.HandleFunc("/fetchdata", h.getData).Methods("POST")

	// admin routes
	// router.HandleFunc("/users/{userID}", auth.WithJWTAuth(h.handleGetUser, h.store)).Methods(http.MethodGet)
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {

	var fd types.PostPayload

	if err := utils.ParseJSON(r, &fd); err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	er := h.post.PostContent(fd)

	if er != nil {
		utils.WriteError(w, http.StatusBadRequest, er)
		log.Println("Error", er)
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "done"})
}


func (h *Handler) getData(w http.ResponseWriter, r *http.Request) {
	log.Println("ok done")

	var mail types.ReqBody


	if err := utils.ParseJSON(r, &mail); err != nil {
		log.Println(r)
		log.Println("mail")
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	s, er := h.post.GetPostedData(mail)

	if er != nil {
		log.Println("Error getting posted data:", er)
		utils.WriteError(w, http.StatusInternalServerError, er)
		return
	}
	log.Println(er)
	log.Println(s)

	// Return the data as a JSON response
	utils.WriteJSON(w, http.StatusOK, s)

}
