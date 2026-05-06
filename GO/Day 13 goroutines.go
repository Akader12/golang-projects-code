// DAY 13 -Goroutines & Channels
package main

// import (
// 	"fmt"
// 	"time"
// )

// 1. Goroutines -> is ligthtweight thread managed by the go runtime

// func sayHello()  {
// 	fmt.Println("Hello from Goroutine")
// }

// func main()  {
// 	go sayHello() //start goroutine
// 	time.Sleep(time.Second)  //wait for it to finish
// 	fmt.Println("main function done")
// }

// 2. Goroutine with parameters
// func printNumber(n int)  {
// 	fmt.Println("number:",n)
// }

// func main()  {
// 	for i := 1;i<=5;i++{
// 		go printNumber(i)
// 	}
// 	time.Sleep(time.Second)
// }

// 3. Channel -> is a wway for goroutines to communicate with each other an synchronize execution

// func sendMessage(ch chan string,)  {
// 	ch <- "hello from channel"
// }
// func main()  {
// 	messageChannel := make(chan string) //create channel
	
// 	go sendMessage(messageChannel)

// 	msg := <-messageChannel //receive from channel
// 	fmt.Println(msg)
	
// }


// channels (basic)

// import "fmt"
// // ch := make(chan int) //create
// // ch <- 10			//send
// // x := <-ch		// receive

// func sendData(ch chan int)  {
// 	ch <- 10
// }

// func main()  {
// 	ch := make(chan int)

// 	go sendData(ch)
// 	data := <- ch
// 	fmt.Println(data)
// }