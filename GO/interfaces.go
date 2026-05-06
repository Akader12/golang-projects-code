package main

// import (
// 	"fmt"
// 	"math"
// )

// type Shape interface {
// 	Area() float64
// 	Circum() float64
// }

// // square struct
// type Square struct {
// 	length float64
// }
// type Circle struct {
// 	radius float64
// }

// // square methods
// func (s Square) Area() float64 {
// 	return s.length * s.length
// }
// func (s Square) Circum() float64 {
// 	return s.length * 4
// }
// //circle methods
// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius *c.radius
// }
// func (c Circle) Circum() float64 {
// 	return 2 * math.Pi * c.radius
// }
// func printInfo(s Shape)  {
// 	fmt.Printf("Area of %T is: %0.2f \n",s,s.Area())
// 	fmt.Printf("Circumference of %T is: %0.2f \n",s,s.Circum())
// }


// func main()  {
// 	shapes := []Shape{
// 		Square{length: 20},
// 		Circle{radius: 40},
// 		Circle{radius: 40},
// 		Square{length: 40},
// 	}

// 	for _,v := range shapes{
// 		printInfo(v)
// 	}
// }