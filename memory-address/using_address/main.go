package main

import "fmt"

// Print: standard output
// Println: standard output with newline
// Printf: standard output according to format specifier
// Scan: standard input

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	// Scan puts the input value into the meters mem location
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
