package main

import "sort"

func main() {
	
}

type RowPower struct {
	 Row int
	 Power int
}

// ByPower is a sort.Interface for []RowPower
// based on the power
type ByPower []RowPower

func (b ByPower) Len() int {
	return len(b)
}
func (b ByPower) Swap(i, j int)  {
	b[i], b[j] = b[j], b[i]
}
func (b ByPower) Less(i, j int) bool {
	return b[i].Power < b[j].Power
}

func kWeakestRows(mat [][]int, k int) []int {
	//rowPower := make(map[int]int, 0)
	// switching to struct, not hashed like a map and can be ordered

	// go through mat and find power of each row
	// O(NM)
	var rps []RowPower
	for k, row := range mat {
		//rp = append(rp, RowPower{Row: k,})
		rp := RowPower{Row: k, }
		for _, v := range row {
			rp.Power = v
		}
		rps = append(rps, rp)
	}

	// slice of rows strength in order
	// O(N) can be optimized to just to k.. maybe
	// return [:k]
	sort.Sort(ByPower(rps))
	return rps[:k]
}