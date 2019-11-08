package main

import "fmt"

func main() {
	NewFoo([]int{1, 2, 3})
}

//Foo public object
type Foo struct {
	order []int
}

func (f *Foo) first(out1 chan<- struct{}) {
	fmt.Println("inside first")
	out1 <- struct{}{}
	fmt.Print("first")
}
func (f *Foo) second(in1 <-chan struct{}, out2 chan<- struct{}) {
	<-in1
	out2 <- struct{}{}
	fmt.Print("second")
}
func (f *Foo) third(in2 <-chan struct{}, out3 chan<- struct{}) {
	<-in2
	out3 <- struct{}{}
	fmt.Print("third")
}

//NewFoo is a constructor for a new object Foo
//takes input of slice of ints
func NewFoo(order []int) *Foo {
	fmt.Println("Creating foo object")
	f := &Foo{order}
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	fmt.Printf("Len: %d\n", len(f.order))
	for _, o := range f.order {
		switch o {
		case 1:
			fmt.Println("creating 1st goroutine")
			go f.first(ch1)
			fmt.Println("first complete")
		case 2:
			fmt.Println("creating 2nd goroutine")
			go f.second(ch1, ch2)
			fmt.Println("second complete")
		case 3:
			fmt.Println("creating 3rd goroutine")
			go f.third(ch2, ch3)
			fmt.Println("tird complete")
		}
	}
	<-ch3
	return f
}
