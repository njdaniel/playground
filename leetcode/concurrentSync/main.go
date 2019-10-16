package main

import "fmt"

func main() {
	
}

//Foo public object
type Foo struct {
	order []int
}
func (f *Foo) first() {
	fmt.Print("first")
}
func (f *Foo) second() {
	fmt.Print("second")
}
func (f *Foo) third() {
	fmt.Print("third")
}
//NewFoo is a constructor for a new object Foo
//takes input of slice of ints
func NewFoo(order []int) *Foo {
	f := &Foo{order}
	for _, o := range f.order {
		switch o {
		case 1:
			go f.first()
		case 2:
			go f.second()
		case 3:
			go f.third()
		}
	}
	return f
}

