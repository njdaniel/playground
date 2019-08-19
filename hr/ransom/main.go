package main

import "fmt"

func main() {
	
}

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	m := make(map[string]int)
	for _, v := range magazine {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	for _, v := range note {
		if _, ok := m[v]; ok && m[v] > 0 {
			m[v] -= 1
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

}