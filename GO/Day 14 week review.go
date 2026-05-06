package main

// import "fmt"

// type Account struct {
// 	Name    string
// 	Balance float64
// }

// // method with pointer receive
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