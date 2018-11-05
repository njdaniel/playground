package main

func main() {

}

//judgeCircle does the robot end up at the start
func judgeCircle(moves string) bool {
	//cmds:
	//  L: left
	//  R: right
	//  U: up
	//  D: down
	//LR and UD cancel each other out
	ms := []rune(moves)
	u := 0
	d := 0
	l := 0
	r := 0
	for _, m := range ms {
		switch m {
		case 'U':
			u++
		case 'D':
			d++
		case 'L':
			l++
		case 'R':
			r++
		}
	}
	if u == d && l == r {
		return true
	}
	return false
}
