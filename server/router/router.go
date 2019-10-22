package router

import (
	"bass/server/app"
	"bass/server/handler"

	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()
	r.Method("GET", "/", handler.NewHandler(a.HandleIndex, l))
	return r
}
