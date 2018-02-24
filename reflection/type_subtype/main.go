package main

import (
	"reflect"
	"fmt"
)

type MyInt int
type IntInt MyInt

func main() {
	var a int
	var b MyInt
	var c IntInt
	refa := reflect.ValueOf(a)
	refb := reflect.ValueOf(b)
	refc := reflect.ValueOf(c)
	fmt.Printf("a \t Type: %T \t Kind: %v \n", a, refa.Kind())
	fmt.Printf("b \t Type: %T \t Kind: %v \n", b, refb.Kind())
	fmt.Printf("b \t Type: %T \t Kind: %v \n", c, refc.Kind())
}
