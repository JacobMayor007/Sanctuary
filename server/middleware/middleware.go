package middleware

import (
	"context"
	"net/http"
	"sanctuary_server/firebase"
	"strings"

	"firebase.google.com/go/v4/auth"
)

type contextKey string

const UserIDKey contextKey = "userID"

// VerifyToken extracts and verifies the Firebase token from the request
func VerifyToken(r *http.Request) (*auth.Token, error) {
	authHeader := r.Header.Get("Authorization")
	idToken := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := firebase.AuthClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authHeader := r.Header.Get("Authorization"); authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Reuse VerifyToken here
		token, err := VerifyToken(r)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
