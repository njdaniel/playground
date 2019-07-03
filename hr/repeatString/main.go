package main

func main() {
	
}

func repeatedString(s string, n int64) int64 {
	count := 0
	d := findA(s)
	sub := findA(s[:(int(n)%(len(s)))])
	count = d * (int(n)/len(s)) + sub
	return int64(count)
}

func findA(s string) int {
	count := 0
	for _, r := range s {
		if r == 'a' {
			count++
		}
	}
	return count
}