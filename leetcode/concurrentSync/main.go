package main

import "fmt"

func main() {
	NewFoo([]int{1,2,3})
}

//Foo public object
type Foo struct {
	order []int
}
func (f *Foo) first(done chan struct{}) {
	fmt.Print("first")
	done <- struct {}{}
}
func (f *Foo) second(done chan struct{}) {
	fmt.Print("second")
	done <- struct{}{}
}
func (f *Foo) third(done chan struct{}) {
	fmt.Print("third")
	done <- struct{}{}
}
//NewFoo is a constructor for a new object Foo
//takes input of slice of ints
func NewFoo(order []int) *Foo {
	fmt.Println("Creating foo object")
	f := &Foo{order}
	done := make(chan struct{})
	fmt.Printf("Len: %d", len(f.order))
	for _, o := range f.order {
		switch o {
		case 1:
			go f.first(done)
		case 2:
			go f.second(done)
		case 3:
			go f.third(done)
		}
	}
	return f
}

