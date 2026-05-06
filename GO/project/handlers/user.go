package handlers

import (
	"encoding/json"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"time"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Name: "ali",
		Age:  20,
	}

	utils.SendSuccesss(w,"User created",user)	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendError(w,"Requst failed","invalid json",400)
		return
	}

	if user.Age == 0 || user.Name == ""{
		utils.SendError(w,"valid error","missing required feilds",400)

		return
	}

	utils.SendSuccesss(w,"user created",user)

}

func HealthCheck(w http.ResponseWriter,r *http.Request)  {
	data := map[string]string{
		"status":"ok",
		"time": time.Now().Format(time.RFC3339),
	}

	utils.SendSuccesss(w,"Server is healthy",data)
}

func Ping(w http.ResponseWriter,r *http.Request)  {
	data := map[string]string{
		"data":"pingo",
		"time":time.Now().Format(time.RFC3339),
	}

	utils.SendSuccesss(w,"pong",data)
}