package main

import "fmt"

// Bit shifting
// KB: KiloBytes 2^10
// MB: MegaBytes 2^20
const (
	_ = iota // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
)

func main() {
	// Printing out binary and decimal of kilobytes and gigabytes
	fmt.Printf("binary\t\tdecimal\n")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}

// iota def is a very small increment value