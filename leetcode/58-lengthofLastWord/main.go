package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello Word"                  //4
	s2 := "   fly me   to   the moon  " //4
	s3 := "luffy is still joyboy"       //6
	fmt.Println(lengthOfLastWord(s1))
	fmt.Println(lengthOfLastWord(s2))
	fmt.Println(lengthOfLastWord(s3))
}

func lengthOfLastWord(s string) int {
	var l int

	//remove leading and trailing whitespace
	s = strings.TrimSpace(s)

	//split s string by " "
	ss := strings.Split(s, " ")
	fmt.Println(ss)
	fmt.Println(ss[len(ss)-1])

	//take len of last word
	l = len(ss[len(ss)-1])

	return l
}
