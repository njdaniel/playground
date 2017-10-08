package main

import (
	"sync"
	"time"
	"math/rand"
	"sync/atomic"
	"fmt"
)

var wg sync.WaitGroup
var counter int64

func main() {
}

func incrementor(s string)  {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) //Access without race
	}
	wg.Done()
}
