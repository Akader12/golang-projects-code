package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func postUser(w http.ResponseWriter, r *http.Request)  {
	
	var user User //declare in main.go

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w,`invalid JSON`,http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w,`user received, name : %v,age: %d`,user.Name,user.Age)

}

func PostRequestmain() {
	r := mux.NewRouter()
	r.HandleFunc("/create-user",postUser)

	fmt.Println(`server running on http://localhost:8080`)
	http.ListenAndServe(`:8080`,r)
}