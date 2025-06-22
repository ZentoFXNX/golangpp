package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= b.Balance {
		b.Balance -= amount
	} else {
		fmt.Println("Недостаточно средств")
	}
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b BankAccount) Speak() {
	fmt.Println("Владелец счета: " + b.Owner)
}

func main() {
	account := BankAccount{Owner: "Bob", Balance: 1000}
	account.Speak()

	account.Deposit(500)
	fmt.Println("Баланс:", account.GetBalance())

	account.Withdraw(300)
	fmt.Println("Баланс:", account.GetBalance())

	account.Withdraw(1500)
	fmt.Println("Баланс:", account.GetBalance())
}