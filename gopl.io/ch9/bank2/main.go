package bank

var (
	sema = make(chan struct{}, 1) // binary semaphore guarding balance
)

func Deposit(amount int) {
	sema <- struct{}{} //acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
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
