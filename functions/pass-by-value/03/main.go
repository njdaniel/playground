package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Todd]
}

func changeMe(z []string){
	z[0] = "Todd"
	fmt.Println(z) // [Todd]
}

// This works because make creates a reference that can be used with slices, channels, and maps
