package main

import (
	"net/http"

	"github.com/andersonlira/album/controller"
)

// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>

func main() {
	mux := http.NewServeMux()
	controller.Register(mux)
	http.ListenAndServe(":7000", mux)
}
