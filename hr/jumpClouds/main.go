package main

func main() {
	
}

func jumpingOnClouds(c []int32) int32 {
	player := 0
	clouds := 0
	var f func(i int)
	f = func(i int) {
		if i+2 > len(c)-1 {
			clouds++
			return
		} else if i+2 == len(c)-1 {
			clouds++
			return
		}
		if c[i+2] == 0 {
			clouds++
			f(i+2)
		} else {
			clouds++
			f(i+1)
		}
	}
	f(player)
	return int32(clouds)
}