package main

import (
	"context"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		hello("John").Render(context.Background(), w)
	})
	http.ListenAndServe(":3003", r)
}
