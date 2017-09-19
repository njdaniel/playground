package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

func (p Person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := Person{"Nick", "Daniel", 27}
	fmt.Println(p1.fullName())
}
