package main

func main() {
	
}

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	m := make(map[rune]int)
	for _, e := range s1 {
		if _, ok := m[e]; !ok {
			m[e] = 1
		}
	}

	for _, e := range s2 {
		if _, ok := m[e]; ok {
			//fmt.Println("YES")
			return "YES"
		}
	}
	//fmt.Println("NO")
	return "NO"
}