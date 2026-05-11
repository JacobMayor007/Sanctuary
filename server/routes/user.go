package routes

import (
	"sanctuary_server/api"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	userApi := api.NewUserApi()

	r.Post("/users/register", userApi.Register)
	r.Get("/users/{userId}", userApi.GetMe)
	r.Put("/users/me", userApi.UpdateMe)
	r.Delete("/users/me", userApi.DeleteMe)
}
