package main

func main() {
	
}

func tribonacci(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 2,
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else  if n == 2 {
		return 1
	}
	sum := 0
	for i := 3; i <= n; i++ {
		if _, ok := m[i-3]; !ok {
			m[i-3] = m[i-5]+m[i-6]+m[i-4]
		}
		if _, ok := m[i-2]; !ok {
			m[i-2] = m[i-5]+m[i-3]+m[i-4]
		}
		if _, ok := m[i-1]; !ok {
			m[i-1] = m[i-2]+m[i-3]+m[i-4]
		}
	}
	sum += m[n-1] + m[n-2] + m[n-3]
	return sum
}