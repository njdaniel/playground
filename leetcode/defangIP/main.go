package main

import "fmt"

func main()  {
	fmt.Println(defangIPaddr("255.100.50.0"))
}

func defangIPaddr(address string) string {
	defIP := make([]rune, 0)
	for _, r := range []rune(address) {
		if r == '.' {
			defIP = append(defIP, '[', '.', ']')
		} else {
			defIP = append(defIP, r)
		}
	}
	return string(defIP)
}