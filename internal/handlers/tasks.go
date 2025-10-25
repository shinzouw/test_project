package handlers

import (
	"awesomeProject1/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tasks)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}

	if task.Title == "" || task.Description == "" {
		http.Error(w, "Title and Description is required", http.StatusBadRequest)
		return
	}
	task.Id = len(models.Tasks) + 1
	task.CreatedAt = time.Now()
	models.Tasks = append(models.Tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Wrong ID", http.StatusBadRequest)
		return
	}
	var NewTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&NewTask); err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}
	for i, task := range models.Tasks {
		if task.Id == id {
			NewTask.Id = id
			NewTask.CreatedAt = task.CreatedAt
			models.Tasks[i] = NewTask
			json.NewEncoder(w).Encode(NewTask)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusBadRequest)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Wrong ID", http.StatusBadRequest)
		return
	}
	for i, task := range models.Tasks {
		if task.Id == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusBadRequest)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Wrong ID", http.StatusBadRequest)
		return
	}
	for _, task := range models.Tasks {
		if task.Id == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
