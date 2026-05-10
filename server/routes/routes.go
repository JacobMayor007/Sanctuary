package routes

import "github.com/go-chi/chi/v5"

func MainRoutes(r chi.Router) {
	UserRoutes(r)
}
