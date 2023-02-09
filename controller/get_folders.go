package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andersonlira/album/service"
)

type FolderHandler struct{}

func (h *FolderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service := service.AwsS3{}
	service.Init()
	service.List()
	json.NewEncoder(w).Encode(service.List())
}

func Register(mux *http.ServeMux) {
	mux.Handle("/folders/", &FolderHandler{})
}
