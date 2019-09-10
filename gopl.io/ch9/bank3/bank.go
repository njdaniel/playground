package bank

import "sync"

var (
	mu      sync.Mutex //guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

//this is very similar to bank2 using mutual exclusion
//this uses sync.Mutex to protect the variables
