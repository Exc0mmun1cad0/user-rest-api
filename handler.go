package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check whether request body in json format
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// Read request body
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate new user ID
	id := uuid.New()

	// Insert new user to database
	mu.Lock()
	db[id.String()] = User{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
		Created:   time.Now(),
	}
	mu.Unlock()

	// Form response with ID of created user
	resp := struct {
		ID string `json:"id"`
	}{
		ID: id.String(),
	}
	rawResp, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(
			w,
			"Invalid request body",
			http.StatusInternalServerError,
		)
		return
	}

	// Send response and set some headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(rawResp)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// Get user from database in case user with this ID exists
	mu.Lock()
	user, exists := db[id]
	mu.Unlock()
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Form response with user data
	resp, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		http.Error(
			w,
			"Failed to encode get user response",
			http.StatusInternalServerError,
		)
		return
	}

	// Send response and set some headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check whether request body in json format
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	id := r.PathValue("id")

	// Check whether user with this ID exists
	mu.Lock()
	newUser, exists := db[id]
	mu.Unlock()
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode request body with new user fields
	var updReq struct {
		FirstName *string `json:"first_name,omitempty"`
		LastName  *string `json:"last_name,omitempty"`
		Email     *string `json:"email,omitempty"`
		Age       *uint   `json:"age,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updReq); err != nil {
		http.Error(
			w,
			"Invalid request body",
			http.StatusInternalServerError,
		)
		return
	}

	// Update newUser with new fields
	if updReq.LastName != nil {
		newUser.LastName = *updReq.LastName
	}
	if updReq.FirstName != nil {
		newUser.FirstName = *updReq.FirstName
	}
	if updReq.Email != nil {
		newUser.Email = *updReq.Email
	}
	if updReq.Age != nil {
		newUser.Age = *updReq.Age
	}

	// Insert newUser in database
	mu.Lock()
	db[id] = newUser
	mu.Unlock()

	// Return updated user in response in order to confirm changes
	resp, err := json.MarshalIndent(newUser, "", "  ")
	if err != nil {
		http.Error(
			w,
			"Failed to encode get user response",
			http.StatusInternalServerError,
		)
		return
	}

	// Send response and set some headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// Check whether user with this id exists
	mu.Lock()
	_, exists := db[id]
	mu.Unlock()
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// User deletion
	mu.Lock()
	delete(db, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
