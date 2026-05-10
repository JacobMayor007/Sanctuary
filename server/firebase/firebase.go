package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client

func InitFirebase() {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Failed to initialize Firebase:", err)
	}

	AuthClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatal("Failed to initialize Firebase Auth:", err)
	}

	log.Println("Firebase initialized!")
}
