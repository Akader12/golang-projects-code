package main

// import "fmt"

// func updateNames(x *string, arr *[3]int) {
// 	*x = "wedge"
// 	arr[0] = 100
// }

// func main() {
// 	name := "tifa"
// 	numbers := [3]int{1,2,3}

// 	updateNames(&name,&numbers)
// 	m := &name
// 	fmt.Println("memory address of name is: ",&name)	
// 	fmt.Println("value at memory address: ",*m)

// 	fmt.Println(name)
// 	fmt.Println(name,numbers)

// }

// type Account struct {
// 	Name    string
// 	Balance float64
// }

// // method with pointer receiver
// func (a *Account) Deposit(amount float64)  {
// 	a.Balance += amount
// }
// func (a *Account) Withdrew(amount float64)  {
// 	if a.Balance >= amount {
// 		a.Balance -= amount
// 	}else {
// 		fmt.Println("not enough")
// 	}
// }

// func (a Account) Display() {
// 	fmt.Printf("Account %s, Balance: $%.2f\n",a.Name,a.Balance)
// }

// func main()  {
// 	acc := Account{Name: "abdikader",Balance: 100}
// 	acc.Display()
// 	acc.Deposit(50)
// 	acc.Display()
// 	acc.Withdrew(150)
// 	acc.Display()
// }