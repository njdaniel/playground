package main

func main() {

}

//fib returns the fibonacci number which is the sum of the last two fibonacci numbers
// constraints 0 <= n <= 30
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fibOpt(n int) int {
	mem := make(map[int]int, 0)
	mem[0] = 0
	mem[1] = 1

	var f func(n int) int
	f = func(n int) int {
		if _, ok := mem[n]; !ok {
			mem[n] = f(n-1) + f(n-2)
		}
		//sum := f(n-1) + f(n-2)
		return mem[n]
	}
	f(n)

	return mem[n]
}
