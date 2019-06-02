package main

import "fmt"

var MaxOutstanding = 3

func main() {
	clientRequests := make(chan *Request, 0)
	stop := make(chan bool, 0)
	go Serve(clientRequests, stop)
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	// Send request
	clientRequests <- request
	// Wait for response.
	fmt.Printf("answer: %d\n", <-request.resultChan)
}

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func Serve(clientRequests chan *Request, quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit
}
