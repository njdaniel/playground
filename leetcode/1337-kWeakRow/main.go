package main

import (
	"fmt"
	"sort"
)

func main() {
	k := 3
	mat := [][]int{{1,1,0,0,0}, {1,1,1,1,0}, {1, 0, 0, 0, 0}, {1,1,0,0,0}, {1,1,1,1,1}}
	fmt.Println(kWeakestRows(mat, k))
}

type RowPower struct {
	 Row int
	 Power int
}

type By func(r1, r2 *RowPower) bool

func (by By) Sort(rows []RowPower)  {
	rs := &RowSorter{
		Rows: rows,
		by:   by,
	}
	sort.Sort(rs)
}

// RowSorter is for using the sort.Sort interface
type RowSorter struct {
	Rows []RowPower
	by func(r1, r2 *RowPower) bool
}
// Len func for sort.Interface
func (r *RowSorter) Len() int {
	return len(r.Rows)
}
// Swap for sort.Interface
func (r *RowSorter) Swap(i, j int)  {
	r.Rows[i], r.Rows[j] = r.Rows[j], r.Rows[i]
}
// Less for sort.Interface
func (r *RowSorter) Less(i, j int) bool {
	return r.by(&r.Rows[i], &r.Rows[j])
}


func kWeakestRows(mat [][]int, k int) []int {
	//rowPower := make(map[int]int, 0)
	// switching to struct, not hashed like a map and can be ordered
	power := func(r1, r2 *RowPower) bool {
		return r1.Power < r2.Power
	}
	// go through mat and find power of each row
	// O(NM)
	var rps []RowPower
	for k, row := range mat {
		//rp = append(rp, RowPower{Row: k,})
		rp := RowPower{Row: k, }
		for _, v := range row {
			if v == 1 {
				rp.Power += 1
			}
		}
		rps = append(rps, rp)
	}

	// slice of rows strength in order
	// O(N) can be optimized to just to k.. maybe
	// return [:k]
	By(power).Sort(rps)
	//return rps[:k]
	fmt.Println(rps)
	weakRows := make([]int, 0)
	// adds O(N)
	for x := 0; x < k; x++ {
		weakRows = append(weakRows, rps[x].Row)
	}

	return weakRows
}