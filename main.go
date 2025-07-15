package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"mini-list/handlers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/mini-list/{archive}/posts.svg", handlers.PostsSvgHandler)

	fmt.Printf("Listening on :8001")
	http.ListenAndServe(":8001", r)
}
