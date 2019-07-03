package main

import "fmt"

func main() {
	var n int32
	n = 7
	ar := []int32{1,2,1,2,1,3,2}
	fmt.Println(sockMerchant(n, ar))
}


func sockMerchant(n int32, ar []int32) int32 {
	var pairs int32
	m := make(map[int32][]int)
	for i := 0; i < int(n); i++ {
		if _, ok := m[ar[i]]; !ok {
			m[ar[i]] = []int{i}
		} else {
			m[ar[i]] = append(m[ar[i]], i)
			if len(m[ar[i]]) % 2 == 0 {
				pairs++
			}
		}

	}

	return pairs
}