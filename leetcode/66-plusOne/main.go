package main

import (
	"fmt"
	"strconv"
)

func main() {
	d := []int{1, 2, 2}
	fmt.Println(d)
}

func plusOne(digits []int) []int {

	//cycle through and convert number to letter(alphanumeric)
	alphanumDigits := ""
	for _, d := range digits {
		alphanumDigits = alphanumDigits + strconv.Itoa(d)
	}
	//fmt.Print(alphanumDigits)

	//convert that into one string

	//convert string to one number
	num, _ := strconv.Atoi(alphanumDigits)

	//add one to number
	num++
	//convert number to string
	alphanumDigits = strconv.Itoa(num)

	//cycle through and create []int
	digitsPlusone := make([]int, 0)
	for _, e := range alphanumDigits {
		d, _ := strconv.Atoi(string(e))
		digitsPlusone = append(digitsPlusone, d)
	}

	//return
	return digitsPlusone
}
