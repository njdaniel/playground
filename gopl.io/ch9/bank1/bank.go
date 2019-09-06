// Package bank provides a concurrency-safe bank with one account.
// This confines the shared variable to one goroutine and brokers the access with channel requests
package bank

var deposits = make(chan int) //send amount to deposit
var balances = make(chan int) //receive balance

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() //monitor goroutine, brokers access to confined variables using channel requests
}
