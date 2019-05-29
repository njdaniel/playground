package main

import (
	"fmt"
	"math"
)

func main() {
	//shapes := []rect{
	//	{3,3,},
	//	{2,2,},
	//}
	shape := rect{3,3}
	ListShapes(shape)
}




type Geometry interface {
	area() float64
	perim() float64
}

type Rect interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func ListShapes(c ...Geometry)  {
	for _, x := range c {
		fmt.Println(x)
	}
}