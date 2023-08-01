package main

import "fmt"

type Account struct {
	AccountNO string
	Pwd string
	Balance float64
}

func (account *Account) Deposit(money float64,pwd string) {
	if pwd != account.Pwd {
		fmt.Println("your pwd is wrong ")
		return
	}

	if money <= 0 {
		fmt.Println("The amount you have entered is not correct")
		return
	}
	account.Balance += money
	fmt.Println("Deposit Success =")
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("your pwd is wrong ")
		return
	}
	fmt.Printf("your balaance is %f =",account.Balance)
}

func main (){
	account := Account{
		AccountNO: "001 ",
		Pwd: "123",
		Balance: 12.3,
	}
	account.Query("123")
}