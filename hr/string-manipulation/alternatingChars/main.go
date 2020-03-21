package main

import "fmt"

func main() {
	s := "AAAABBB"
	fmt.Println(alternatingCharacters(s))
}

// alternatingCharacters will take s and sort into alternating chars of A and B only then delete the min
// chars and return the amount deleted\
// how is this string man? just need to count A and B?
// -- nvm you cant move the chars.. still how is this string manip?
// i guess i still have to delete the string, even though its a copy? passed by value
func alternatingCharacters(s string) int32 {
	var dels int32 //nil value = 0
	//check len
	if len(s) < 2 {
		return dels
	}
	ra := []rune(s)
	for i := 1; i < len(ra); {
		if ra[i] == ra[i-1] {
			ra = append(ra[:i], ra[i+1:]...)
			dels++
		} else {
			i++
		}
	}
	// loop throu and check the prev char
	// if it matches, pop it out, decrease the value of the i
	return dels
}