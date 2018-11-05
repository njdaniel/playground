package main

import "fmt"

func main() {
	myFriend := "Mar"

	switch {
	case len(myFriend) == 3:
		fmt.Println("Your friend only has 3 letters in its name")
	case myFriend == "Mar":
		fmt.Println("and your name is Mar")
	default:
		fmt.Println("No friends")
	}
}

/*
No expression needed for switch just go by case by case
only first one that match prints
*/
