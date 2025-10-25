package handlers

import (
	"net/http"
)

func RegUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
