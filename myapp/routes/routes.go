package routes

import (
	"myapp/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/user", handlers.CreateUser)
	http.HandleFunc("/user/get", handlers.GetUser)
	http.HandleFunc("/user/update", handlers.UpdateUser)
	http.HandleFunc("/user/delete", handlers.DeleteUser)
}
