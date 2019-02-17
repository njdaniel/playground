package main

func main() {
	s := []byte{'h','e', 'l', 'l', 'o'}
	reverseString(s)
	
}

func reverseString(s []byte)  {
	rs := make([]byte, 5)
	for i := len(s)-1; i == 0; i-- {
		rs = append(rs, s[i])
	}
}
