// Package bank provides a concurrency-safe bank with one account.
// This confines the shared variable to one goroutine and brokers the access with channel requests
package bank

import "golang.org/x/text/currency"

var deposits = make(chan int)  //send amount to deposit
var balances = make(chan int)  //receive balance
var withdraws = make(chan int) //send amount to withdraw

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	x := <-balances
	if x >= amount {
		withdraws <- amount
		return true
	}
	return false
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			balance -= amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() //monitor goroutine, brokers access to confined variables using channel requests
}

//ex9.1 add Withdraw func
