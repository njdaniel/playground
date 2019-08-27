package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
	}()

	//squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//printer
	for {
		fmt.Println(<-squares)
	}
	//fatal error: all goroutines are asleep
	
}
