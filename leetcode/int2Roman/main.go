package main

import fmt

func main() {
	i := 30 	 
	fmt.Println(intToRoman(i))
}

func intToRoman(num int) string {
	m := make(map[int]rune{
		1: 'I',
		5: 'V', 
		10: 'X', 
		50: 'L',
		100: 'C',
		500: 'D',
		1000: 'M',
	})

	r := make([]rune, 0)
	n := num
	for ;n>0; {
		for ; n/1000 > 0; {
			n = n-1000
			r = append(r, m[1000] )
		} 
		if n/900 >0 {
			n = n - 900
			r = append(r, []rune{m[100], m[1000]})
		}
		if n/500 >0 {
			n = n - 500
			r = append(r, []rune{m[100], m[1000]})
		}
		if n/400 >0 {
			n = n - 400
			r = append(r, []rune{m[100], m[1000]})
		}
		for ; n/100 > 0; {
			n = n-100
			r = append(r, m[100] )
		} 
		if n/90 >0 {
			n = n - 90
			r = append(r, []rune{m[10], m[100]})
		}
		if n/50 >0 {
			n = n - 50
			r = append(r, m[50])
		}
		if n/40 >0 {
			n = n - 40
			r = append(r, []rune{m[10], m[50]})
		}
		for ; n/10 > 0; {
			n = n-10
			r = append(r, m[10] )
		} 
		if n/9 >0 {
			n = n - 9
			r = append(r, []rune{m[1], m[10]})
		}
		if n/5 >0 {
			n = n - 5
			r = append(r, m[5])
		}
		if n/4 >0 {
			n = n - 4
			r = append(r, []rune{m[1], m[5]})
		}
		for ; n/1 > 0; {
			n = n-1
			r = append(r, m[1] )
		} 
 	}	
	return string(ro)
}
