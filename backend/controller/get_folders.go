package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andersonlira/album/service"
)

type FolderHandler struct{}

func (h *FolderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service := service.AwsS3{}
	service.Init()
	service.List()
	aux := []string{
		"042016/2014-04-17 19.49.24.jpg",
	}
	log.Println(service.GetPreSignedUrls(aux))
	json.NewEncoder(w).Encode(service.List())
}

func Register(mux *http.ServeMux) {
	mux.Handle("/folders/", &FolderHandler{})
}
