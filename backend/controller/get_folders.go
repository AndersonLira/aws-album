package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andersonlira/album/service"
)

type FolderHandler struct{}

func (h *FolderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	service := service.AwsS3{}
	service.Init()
	service.List()
	aux := []string{
		"042016/2014-04-17 19.49.24.jpg",
	}
	log.Println(service.GetPreSignedUrls(aux))
	json.NewEncoder(w).Encode(service.List())
}

type PreSignedUrlsHandler struct{}

func (h *PreSignedUrlsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	values := r.Form["values"]

	service := service.AwsS3{}
	service.Init()
	json.NewEncoder(w).Encode(service.GetPreSignedUrls(values))
}

func Register(mux *http.ServeMux) {
	mux.Handle("/folders/", &FolderHandler{})
	mux.Handle("/urls", &PreSignedUrlsHandler{})
}
