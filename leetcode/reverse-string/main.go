package main

func main() {
	s := []byte{'h','e', 'l', 'y', 'o'}
	reverseString(s)
	
}

func reverseString(s []byte)  {
	//rs := make([]byte, 5)
	//var tmp byte
	//for i := 0; i < len(s); i++{
	//	tmp := s[i]
	//	fmt.Printf("tmp is %v \n", string(tmp))
	//	fmt.Println(string(s[:len(s)-i-1]))
	//}
	for i := len(s)/2-1; i >= 0; i-- {
		opp := len(s)-1-i
		s[i], s[opp] = s[opp], s[i]
	}
}
