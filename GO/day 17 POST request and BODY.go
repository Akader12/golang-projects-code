package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// //POST request
// //Parse JSON body

// type User struct{
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }

// type Product struct{
// 	Name string `json:"name"`
// 	Price float64 `json:"price"`
// }

// func createUser(w http.ResponseWriter, r *http.Request)  {

// 	if r.Method != http.MethodPost{
// 		http.Error(w,"only POST allowed",http.StatusMethodNotAllowed)
// 		return
// 	}
// 	var user User 
	
// 	err := json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil{
// 		http.Error(w,"invalid JSON",400)
// 		return
// 	}

// 	if user.Name == "" || user.Age == 0{
// 		http.Error(w,"missing required fields",http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Name",user.Name)
// 	fmt.Println("Age",user.Age)

// 	//response
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(user)

// }

// func createProduct(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost{
// 		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
// 		return
// 	}
// 	var product Product

// 	decoder := json.NewDecoder(r.Body)
// 	decoder.DisallowUnknownFields()
	
// 	err := decoder.Decode(&product)
// 	if err != nil {
// 		fmt.Fprintf(w,"invalid json",http.StatusBadRequest)
// 		return
// 	}
	
// 	//printing
// 	fmt.Println("Name:",product.Name)
// 	fmt.Println("Price:",product.Price)

// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(product)


// }

// func main()  {
// 	http.HandleFunc("/user",createUser)
// 	http.HandleFunc("/product",createProduct)

// 	http.ListenAndServe(":8080",nil)
// }

