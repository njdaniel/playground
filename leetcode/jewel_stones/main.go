package main

func main() {
	
}

func numJewelsInStones(J string, S string) int {
	num := 0
	for _, s := range S {
		for _, j := range J {
			if s == j {
				num++
			}
		}
	}
	return num
}
