package cat

import (
	"encoding/json"
	"meowbase/internal/helpers"
	"net/http"

	"github.com/go-chi/chi"
)

type CatController struct{}

func (cc CatController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", cc.List)
	r.Post("/", cc.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", cc.Get)
		r.Put("/", cc.Update)
		r.Delete("/", cc.Delete)
	})

	return r
}

func (cc CatController) List(w http.ResponseWriter, r *http.Request) {
	c := GetCats()
	resp := &helpers.Response{Writer: w}
	resp.Status = http.StatusOK
	resp.Data = c
	resp.Respond()
}

func (cc CatController) Create(w http.ResponseWriter, r *http.Request) {
	c := &Cat{}
	resp := &helpers.Response{Writer: w}

	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		panic(err)
	}

	err = c.Add()
	if err != nil {
		panic(err)
	}

	resp.Status = http.StatusOK
	resp.Data = c
	resp.Respond()
}

func (cc CatController) Get(w http.ResponseWriter, r *http.Request) {
	c := &Cat{}
	resp := &helpers.Response{Writer: w}
	catID := helpers.GetIDFromRequest(r)

	err := c.Get(catID)
	if err != nil {
		panic(err)
	}

	resp.Status = http.StatusOK
	resp.Data = c
	resp.Respond()
}

func (cc CatController) Update(w http.ResponseWriter, r *http.Request) {
	c := &Cat{}
	resp := &helpers.Response{Writer: w}
	catID := helpers.GetIDFromRequest(r)

	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		panic(err)
	}

	err = c.Update(catID)
	if err != nil {
		panic(err)
	}

	resp.Status = http.StatusOK
	resp.Data = c
	resp.Respond()

}

func (cc CatController) Delete(w http.ResponseWriter, r *http.Request) {
	c := &Cat{}
	resp := &helpers.Response{Writer: w}
	catID := helpers.GetIDFromRequest(r)

	err := c.Delete(catID)
	if err != nil {
		panic(err)
	}

	resp.Status = http.StatusOK
	resp.Respond()
}
