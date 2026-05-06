package handlers

import (
	"encoding/json"
	"fmt"
	"myapp/database"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"time"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows,err := database.DB.Query(r.Context(),"SELECT id,name,age FROM users")
	if err != nil {
		utils.SendError(w,"Query failed",err.Error(),500)
		return
	}

	defer rows.Close()
	
	var users []models.User
	for rows.Next(){
		var user models.User
		err := rows.Scan(&user.ID,&user.Name,&user.Age)
		if err != nil {
			utils.SendError(w,"scan failed",err.Error(),http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		utils.SendError(w,"",err.Error(),500)
		return
	}

	fmt.Println(users)
	utils.SendSuccesss(w,"users seleted",users)
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	id := r.URL.Query().Get("id")
	var user models.User
	err := database.DB.QueryRow(r.Context(),"SELECT id,name,age FROM users WHERE id=$1",id).Scan(&user.ID,&user.Name,&user.Age)
	if err != nil {
		utils.SendError(w,"","user not found",500)
		return
	}

	utils.SendSuccesss(w,"user selected",user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	id := r.URL.Query().Get("id")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	_,err := database.DB.Exec(r.Context(),"UPDATE users SET id=$1, name=$2, age=$3 WHERE id=$4",user.ID,user.Name,user.Age,id)
	if err != nil {
		utils.SendError(w,"Update failed",err.Error(),500)
		return
	}
	utils.SendSuccesss(w,"user updated",user)
}


func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	id :=  r.URL.Query().Get("id")
	_,err := database.DB.Exec(r.Context(),"DELETE FROM users WHERE id=$1",id)

	if err != nil {
		utils.SendError(w,"delete failed",err.Error(),500)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":"user deleted",
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err :=json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.SendError(w,"invalid json",err.Error(),http.StatusInternalServerError)
		return
	}
	
	if user.Name == "" || user.Age ==0{
		utils.SendError(w,"","missing required feild",500)
		return
	}
	_,err = database.DB.Exec(r.Context(),"INSERT INTO users(name,age) VALUES($1,$2)",user.Name,user.Age)
	if err != nil {
		utils.SendError(w,"database failed",err.Error(),500)
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

