package main

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// func main() {
// 	//strings package

// 	greeting := "hello there freinds"

// 	fmt.Println(strings.Contains(greeting,"hello?"))
// 	fmt.Println(strings.ReplaceAll(greeting,"hello","hi"))
// 	fmt.Println(strings.ToUpper(greeting))
// 	fmt.Println(strings.Index(greeting,"ll")) 
// 	fmt.Println(strings.Split(greeting," "))
	
// 	//oringinal value is unchanged
// 	fmt.Println("original string value :",greeting)

// 	ages := []int{45,20,35,30,75,60,25}

// 	sort.Ints(ages) //-> upates the original value outomatically
// 	fmt.Println(ages)

// 	index := sort.SearchInts(ages,300) // The return value is the index to insert x if x is not present (it could be len(ages))
// 	fmt.Println(index)
	
// 	names := []string{"Ak","ad","er","mo"}

// 	sort.Strings(names)
// 	fmt.Println(names)

// 	fmt.Println(sort.SearchStrings(names,"Ak"))
// }