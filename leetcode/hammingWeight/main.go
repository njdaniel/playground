package main

import (
	"fmt"
	"strconv"
)

func main() {
}

func hammingWeight(num uint32) int {
	//TODO: will this verify base2 number?
	fmt.Printf("input %b", num)
	//buf := make([]byte, binary.MaxVarintLen32)
	//B := binary.PutUvarint(buf, uint64((num)))
	//fmt.Printf("binary: %b", buf[:B])
	sum := 0
	numstr := strconv.FormatUint(uint64(num), 2)
	for _, v := range numstr {
		if v == '1' {
			sum++
		}
	}
	return sum
}

//if called many times, use a closure to store the num as a key
func hammingWeightOpt(num uint32) int {
	m := make(map[uint32]int)
	var f = func() int {
		if _, ok := m[num]; ok {
			return m[num]
		}
		sum := 0
		numstr := strconv.FormatUint(uint64(num), 2)
		for _, v := range numstr {
			if v == '1' {
				sum++
			}
		}
		m[num] = sum
		return sum
	}
	return f()
}
