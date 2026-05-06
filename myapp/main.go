package main

import (
	"context"
	"fmt"
	"myapp/database"
	"myapp/routes"
	"net/http"
)

func main()  {
	database.ConnectB()

	defer database.DB.Close(context.Background())

	routes.SetupRoutes()
	http.ListenAndServe(":8080",nil)


	fmt.Println("server running on http://localhost:8080")

}