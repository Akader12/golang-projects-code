package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// code		meaning
// 200		OK(success)
// 201		Created(POST success)
// 400		Bad Request(user error)
// 404		Not Found
// 500		Server Error

// func getUserBYid(w http.ResponseWriter, r *http.Request)  {
// 	id := r.URL.Query().Get("id")

// 	if id == ""{
// 		http.Error(w,"id is required",400)
// 		return
// 	}

// 	w.Header().Set("Content-Type","applcation/json")
// 	w.WriteHeader(200)

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"id":id,
// 	})
// }

// func getUserByUserName(w http.ResponseWriter, r *http.Request)  {
// 	username := r.URL.Query().Get("username")

// 	if username == ""{
// 		http.Error(w,"username is required",400)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"username":username,
// 	})
// }

// func main()  {
// 	http.HandleFunc("/search",getUserBYid)
// 	http.HandleFunc("/username",getUserByUserName)

// 	http.ListenAndServe(":8080",nil)
// }