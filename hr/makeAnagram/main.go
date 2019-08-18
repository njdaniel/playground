package main

func main() {
	
}

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	d := 0
	m := make(map[rune]int)
	//put a in map[rune] num int
	for _, r := range a {
		if _, ok := m[r]; !ok {
			m[r] = 1
		} else {
			m[r] += 1
		}
	}

	//loop b, lookup a if not +1 to return if yes -1 to num
	for _, r := range b {
		if _, ok := m[r]; !ok {
			d += 1
		} else {
			m[r] -= 1
			if m[r] < 1 {
				delete(m, r)
			}
		}
	}

	//loop a and add nums to return
	for _, v := range m {
		d += v
	}

	return int32(d)
}