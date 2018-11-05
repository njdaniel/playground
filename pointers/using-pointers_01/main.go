package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)

	*b = 42        // b says, "The value at this address is changed to 42
	fmt.Println(a) // 42

	// can pass mem addresses instead of values
	// and can stil change the value of whats beings stored
	// better performant
	// no need to pass large amounts of data
}
