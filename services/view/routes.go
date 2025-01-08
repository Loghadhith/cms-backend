package view

import "github.com/Loghadhith/cms/types"

type Handler struct {
	view types.ViewStore
}

func NewHandler(view types.ViewStore) *Handler {
	return &Handler{view: view}
}
