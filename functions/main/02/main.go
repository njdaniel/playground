package main

import "fmt"

func main() {
	greet("Nick", "Daniel")
}

//funct gree(fname string, lname string) {
func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
