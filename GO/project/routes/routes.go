package routes

import (
	"myapp/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/user", handlers.GetUser)
	http.HandleFunc("/user/create", handlers.CreateUser)
	http.HandleFunc("/healthy", handlers.HealthCheck)
	http.HandleFunc("/ping", handlers.Ping)
}
