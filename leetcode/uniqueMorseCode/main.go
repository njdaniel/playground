package main

import "fmt"

func main() {
	words := []string{
		"gin",
		"zen",
		"gig",
		"msg",
	}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {
	m := map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}
	um := make(map[string]string, 0)
	for _, w := range words {
		mw := ""
		for _, r := range []rune(w) {
			mw = fmt.Sprintf("%s%s",mw,m[r])
		}
		//check unique morses
		if _, ok := um[mw]; !ok {
			um[mw] = mw
		}
	}
	fmt.Println(um)
	return len(um)
}
