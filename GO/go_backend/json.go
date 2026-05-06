package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type User struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// 	Job string `json:"job"`
// }

// func createUser(w http.ResponseWriter, r *http.Request)  {
// 	user := User{
// 		Name: `biite`,
// 		Age: 12,
// 		Job: `developer`,
// 	}
// 	w.Header().Set(`content-type`,`application/json`)
// 	json.NewEncoder(w).Encode(user)
// }

// func main() {
// 	http.HandleFunc(`/`,createUser)

// 	http.ListenAndServe(`:8080`,nil)
// }