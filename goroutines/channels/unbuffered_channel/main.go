package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <10; i++{
			c <- i
		}
	}()

	go func() {
		fmt.Println(<-c)
	}()

	time.Sleep(time.Second)
}
