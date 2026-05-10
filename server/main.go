package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sanctuary_server/database"
	"sanctuary_server/firebase"
	"sanctuary_server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	r := chi.NewRouter()
	database.ConnectDatabase()
	firebase.InitFirebase()
	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	routes.MainRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println(`
	===================================================
		Welcome to the Sanctuary Server
	===================================================
	`)

	log.Printf("Sanctuary running on :%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
