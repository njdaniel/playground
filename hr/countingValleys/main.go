package main

import "fmt"

func main() {
	var n int32
	n = 8
	s := "DDUUUUDD"
	fmt.Println(countingValleys(n, s))
}

func countingValleys(n int32, s string) int32 {
	var valleys int32
	lvl := 0
	if len(s) == 0 {
		return 0
	}

	if s[0] == 'D' {
		valleys++
		lvl--
	} else if s[0] == 'U' {
			lvl++
	}

	for i := 1; i < int(n); i++ {
		if lvl == 0 && s[i] == 'D' {
			valleys++
			lvl--
		} else if s[i] == 'D' {
				lvl--
		} else if s[i] == 'U' {
					lvl++
		}
	}

	return valleys
}