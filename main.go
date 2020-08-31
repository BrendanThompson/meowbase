package main

import (
	"meowbase/internal/cat"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("v0.0.0"))
	})

	r.Mount("/cats", cat.CatController{}.Routes())

	err := http.ListenAndServe(":80", r)
	if err != nil {
		panic(err)
	}
}
