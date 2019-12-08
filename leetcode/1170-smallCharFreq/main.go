package main

func main() {
	
}

func numSmallerByFrequency(queries []string, words []string) []int {
	// put data into data structure to include the len of the smallest char
	// data []int: for each query include the number of greater than the words
	// O(QWL)

	ss := make([]int, 0)
	//func for finding the smallest UTF-8 char: O(L)
	var fsf = func(wq string) int {
		var sc rune
		f := 0
		for i, v := range wq {
			if i == 0 {
				sc = v
				f = 1
			} else if v < sc {
				sc = v
				f = 1
			} else if v == sc {
				f++
			}
		}
		return f
	}

	for i, q := range queries {
		qf := fsf(q)
		ss = append(ss, 0)
		for _, w := range words {
			wf := fsf(w)
			if qf < wf {
				ss[i]++
			}
		}
	}

	// for each query and word need to find the smallest char
	return ss
}