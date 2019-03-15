package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nput := []string{"9001 discuss.leetcode.com"}
	fmt.Println(subdomainVisits(nput))
}

func subdomainVisits(cpdomains []string) []string {
	//init return slice string use map for repeats
	subdomains := make(map[string]int, 0)
	//loop thro slice
	for _, v := range cpdomains {
		// split the slice from domain and int
		sp := strings.Fields(v)
		// check if in map if not add otherwise update the int
		// if _, ok := subdomains[sp[1]]; !ok {
		// 	visits, _ := strconv.Atoi(sp[0])
		// 	subdomains[sp[1]] = visits
		// } else {
		// 	visits, _ := strconv.Atoi(sp[0])
		// 	subdomains[sp[1]] = subdomains[sp[1]] + visits
		// }
		// get num of .'s
		num := 0
		for j := 0; j < len(sp[1]); j++ {
			if sp[1][j] == '.' {
				num++
			}
		}
		sepDomains := strings.Split(sp[1], ".")
		// loop thro for each subdomain
		for i := num; i >= 0; i-- {
			curDomain := strings.Join(sepDomains[i:], ".")
			visits, _ := strconv.Atoi(sp[0])
			subdomains[curDomain] += visits
		}

	}
	ss := make([]string, 0)
	for k, v := range subdomains {
		s := fmt.Sprintf("%v %v", strconv.Itoa(v), k)
		ss = append(ss, s)
	}
	return ss
}
