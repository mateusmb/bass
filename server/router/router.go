package router

import (
	"bass/server/app"
	"bass/server/handler"

	"bass/server/router/middleware"

	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()

	// Routes for healthz
	r.Get("/healthz/liveness", app.HandleLive)
	r.Method("GET", "/healthz/readiness", handler.NewHandler(a.HandleReady, l))

	r.Method("GET", "/", handler.NewHandler(a.HandleIndex, l))

	// Routes for APIs
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)

		// Routes for bookings
		r.Method("GET", "/bookings", handler.NewHandler(a.HandleListBookings, l))
		r.Method("POST", "/bookings", handler.NewHandler(a.HandleCreateBooking, l))
		r.Method("GET", "/bookings/{id}", handler.NewHandler(a.HandleReadBooking, l))
		r.Method("PUT", "/bookings/{id}", handler.NewHandler(a.HandleUpdateBooking, l))
		r.Method("DELETE", "/bookings/{id}", handler.NewHandler(a.HandleDeleteBooking, l))
	})

	return r
}
