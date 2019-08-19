package mistake

import "fmt"

// common mistake, since the go routines are called but nothing make it wait for them to finish
// nor does this print in the order Foobar that is desired
func main() {
	for i := 0; i < 10; i++ {
		go foo()
		go bar()
	}
}

func foo()  {
	fmt.Printf("Foo")
}

func bar()  {
	fmt.Printf("bar \n")
}


