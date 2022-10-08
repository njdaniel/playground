package main

func main() {

}

func largestUniqueNumber(nums []int) int {
	lUN := -1
	//brute force
	//eliminate dupes
	unique := make(map[int]int, 0)
	for k := range nums {
		if _, ok := unique[k]; !ok {
			unique[k] = 1
		} else {
			unique[k] = 0
		}
	}
	//find biggest numb
	for k, v := range unique {
		if k > lUN && v == 1 {
			lUN = k
		}
	}

	return lUN
}
