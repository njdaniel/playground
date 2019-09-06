// Package bank provides a concurrency-safe bank with one account.
// This confines the shared variable to one goroutine and brokers the access with channel requests
package bank

type Withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)           //send amount to deposit
var balances = make(chan int)           //receive balance
var withdrawals = make(chan Withdrawal) //send amount to withdraw

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- Withdrawal{amount, ch}
	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawals:
			if w.amount > balance {
				w.success <- false
				continue
			}
			balance -= w.amount
			w.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() //monitor goroutine, brokers access to confined variables using channel requests
}

//ex9.1 add Withdraw func
