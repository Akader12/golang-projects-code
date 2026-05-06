package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // //	Query params(?id=1)
// // //	URL params(/user/1)
// // //  how to read them

// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func search(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name")
// 	age := r.URL.Query().Get("age")

// 	fmt.Println("Name:", name)
// 	fmt.Println("Age:", age)

// 	fmt.Fprintln(w, "query received")
// }

// func getUser(c *gin.Context) {
// 	id := c.Param("id")
// 	postId := c.Param("postId")

// 	c.JSON(http.StatusOK, gin.H{
// 		"user_id": id,
// 		"post_id": postId,
// 	})
// }

// func getProduct(c *gin.Context) {
// 	id := c.Param("id")

// 	c.JSON(http.StatusOK, gin.H{
// 		"product_id": id,
// 		"status":     "ok",
// 	})
// }

// func sendError(c *gin.Context, err string, code int) {
// 	c.JSON(code, gin.H{
// 		"error": err,
// 	})
// }

// func getRides(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		sendError(c,"id must be a number",400)
// 		return
// 	}
// 	if id != 1 {		
// 		sendError(c,"ride not found",400)
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"error": "active ride",
// 	})

// }

// func main() {
// 	r := gin.Default()

// 	//route with URL parameter
// 	r.GET("/users/:id/posts/:postId", getUser)
// 	r.GET("/product/:id", getProduct)
// 	r.GET("/rides/:id", getRides)
// 	http.HandleFunc("/search", search)

// 	r.Run(":8080")
// }
