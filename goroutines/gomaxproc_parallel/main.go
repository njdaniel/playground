package main

import (
	"runtime"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

}

func foo()  {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar()  {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}