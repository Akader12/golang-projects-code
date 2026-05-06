package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// func userIdHandler(w http.ResponseWriter , r *http.Request)  {
// 	vars := mux.Vars(r);
// 	id := vars[`id`];

// 	fmt.Fprintf(w,`user ID.: %s`,id)
// }

// func userHandler(w http.ResponseWriter, r *http.Request){
// 	user := &User{
// 		Id: 1,
// 		Name: `akader`,
// 		Age: 23,
// 	}

// 	if err := json.NewEncoder(w).Encode(user);err != nil{
// 		http.Error(w,`couldn't encode user`,http.StatusBadRequest)
// 		return
// 	}
// }

func main() {

	r := mux.NewRouter()

	// r.HandleFunc(`/users/`,userHandler).Methods(`GET`)
	// r.HandleFunc("/users/{id}",userIdHandler).Methods("GET")

	//middlewares
	r.Handle(`/logmiddleware`, LogMiddleware(http.HandlerFunc(HelloHandler)))
	r.Handle(`/errormiddleware`, ErrorMiddleware(http.HandlerFunc(HelloHandler)))
	r.Handle(`/authmiddleware`, AuthMiddleware(http.HandlerFunc(HelloHandler)))

	//combined middleware: Log -> Auth -> Error -> Handler
	finalHandler := LogMiddleware(AuthMiddleware(ErrorMiddleware(http.HandlerFunc(HelloHandler))))

	r.Handle("/",finalHandler)
	fmt.Println("server running on :8080")

	http.ListenAndServe(":8080", r)

}
