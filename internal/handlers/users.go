package handlers

import (
	"awesomeProject1/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Wrong ID", http.StatusBadRequest)
		return
	}
	for _, user := range models.Users {
		if user.Id == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Wrong ID", http.StatusBadRequest)
		return
	}
	var NewUser models.User
	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	for i, user := range models.Users {
		if user.Id == id {
			NewUser.Id = id
			NewUser.CreatedAt = user.CreatedAt
			models.Users[i] = NewUser
			json.NewEncoder(w).Encode(NewUser)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
