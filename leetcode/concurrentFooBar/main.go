package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// printing foo and bar in order with goroutines
func main() {
	fmt.Println("starting..")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go bar()
		go foo()
		wg.Wait()
	}

}

func foo()  {
	fmt.Printf("foo")
	wg.Done()
}

func bar()  {
	fmt.Printf("bar\n")
	wg.Done()
}
