package main

import "fmt"

func main() {
	fmt.Println(greet("Nick", "Daniel"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, " ", lname)
}
