package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesusandres31/go-gorm-restapi/db"
	"github.com/jesusandres31/go-gorm-restapi/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
}
