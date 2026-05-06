package models

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

// type Response struct{
// 	Message string `json:"message"`
// 	Data any `json:"data,omitempty"`
// }

// type ErrResponse struct{
// 	Message string `json:"message"`
// 	Error string `json:"error"`
// }