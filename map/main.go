package main

import "fmt"

func main() {
	m := make(map[string][]int)
	m["tiger"] = []int{3}
	t := "lion"
	if _, ok := m[t]; ok {
		fmt.Println(m[t])
	} else {
		m[t] = append(m[t], 4)
		fmt.Println(m[t])
	}

	for _, v := range m[t] {
		switch v {
		case 4:
			fmt.Println("just 4")
		default:
			fmt.Println("sjhit")

		}
	}

	for k, meep := range m {
		fmt.Printf("%s: ",k)
		fmt.Println(meep)
	}
}
