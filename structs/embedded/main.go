package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) fullName() string  {
	return p.first + p.last
}

func main() {
	p1 := DoubleZero{
		Person: Person{"Nick", "Daniel", 27},
		LicenseToKill: true,
	}
	fmt.Println(p1.fullName())
	fmt.Println(p1.LicenseToKill)
}
