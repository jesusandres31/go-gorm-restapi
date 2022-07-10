package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesusandres31/go-gorm-restapi/db"
	"github.com/jesusandres31/go-gorm-restapi/models"
	"github.com/jesusandres31/go-gorm-restapi/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// database connection
	db.DBConnection()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	// Index route
	r.HandleFunc("/", routes.HomeHandler)

	s := r.PathPrefix("/api").Subrouter()

	// tasks routes
	s.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	s.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	s.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")

	// users routes
	s.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	s.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	s.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	s.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":4000", r)

}
