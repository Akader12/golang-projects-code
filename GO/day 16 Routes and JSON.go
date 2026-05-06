package main

//Routes
//Serve JSON

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// // 1. Multiple routes
// func home(w http.ResponseWriter,r *http.Request) {
// 	fmt.Fprintf(w,"Home Page")
// }

// func about(w http.ResponseWriter,r *http.Request) {
// 	fmt.Fprintf(w,"About Page")
// }

// // 2. JSON Response

// type User struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }

// type Product struct {
// 	Name string `json:"name"`
// 	Price float64 `json:"price"`
// }

// func getUser(w http.ResponseWriter,r *http.Request)  {
// 	user := User{
// 		Name: "ali",
// 		Age: 23,
// 	}
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(user)
// }

// func getProduct(w http.ResponseWriter,r *http.Request)  {
// 	product := Product{
// 		Name: "phone",
// 		Price: 150.0,
// 	}
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func main() {
// 	http.HandleFunc("/home",home)
// 	http.HandleFunc("/about",about)
// 	http.HandleFunc("/user",getUser)
// 	http.HandleFunc("/product",getProduct)



// 	http.ListenAndServe(":8080",nil)
// }


