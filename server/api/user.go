package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sanctuary_server/middleware"
	"sanctuary_server/repository"

	"github.com/go-chi/chi/v5"
)

type UserApi struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserApi() *UserApi {
	return &UserApi{
		UserRepository: repository.NewUserRepository(),
	}
}

// POST /users/register
func (u *UserApi) Register(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.VerifyToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	uid := token.UID
	email, _ := token.Claims["email"].(string)

	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := u.UserRepository.CreateUser(uid, email, body.Name)
	if err != nil {
		if err.Error() == "user already exists" {
			http.Error(w, "User already registered", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GET /users/me
func (u *UserApi) GetMe(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.VerifyToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 1. Get the parameter as a string
	idStr := chi.URLParam(r, "userId")

	// 2. Convert string to int using strconv.Atoi
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	user, err := u.UserRepository.GetUserByID(uint(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// PUT /users/me
func (u *UserApi) UpdateMe(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.VerifyToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := u.UserRepository.UpdateUser(token.UID, body.Name)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DELETE /users/me
func (u *UserApi) DeleteMe(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.VerifyToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = u.UserRepository.DeleteUser(token.UID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
