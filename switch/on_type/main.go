package main

import "fmt"

// switch base on types
// normally switch on values
//  this allows you to switch base on the type of the variable

type Contact struct {
	greeting string
	name string
}

func SwitchOnType(x interface{})  {
	switch x.(type) {	// This is an assert; asserting "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}

func main() {

}
