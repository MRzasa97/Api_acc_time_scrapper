package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux, env *Env) {

	r.Route("/acc", func(r chi.Router) {
		r.Use(middleware.StripSlashes)
		r.Post("/create", env.CreateRecord)
		r.Get("/records", env.GetAllRecords)
	})
}
