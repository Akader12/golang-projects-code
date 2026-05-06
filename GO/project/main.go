package main

import (
	"fmt"
	"net/http"
	"myapp/routes"
)

func main()  {
	routes.SetupRoutes()

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}