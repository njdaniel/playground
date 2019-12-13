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

// words only need to be called once O(Q+W) not O(QW)
func numSmallerByFrequencyOpt(queries []string, words []string) []int {
	// call fsf func for each and put into a []int
	ss := make([]int, 0)
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
	ws := make([]int, 0)
	for _, word := range words {
		ws = append(ws, fsf(word))
	}
	
	for _, query := range queries {
		qf := fsf(query)
		// do binary search
		ss = append(ss, qf)
	}
	

	// search with binary search
	
	
	return ss
}

func binarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		result += mid + 1
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}

