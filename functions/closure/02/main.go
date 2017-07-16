package main

import "fmt"

var x = 0
//var x int

func increment() int  {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for 2 or more funcs to have to access to the same variable,
that variable would need to be package scope
 */
