package addbinary

import (
	"strconv"
)

func addBinary(a, b string) string {
	// ai, _ := strconv.Atoi(a)
	// bi, _ := strconv.Atoi(b)
	// ab, _ := strconv.ParseInt(a, 2, 64)
	// bb, _ := strconv.ParseInt(b, 2, 64)
	// ab := strconv.ParseInt()
	ai, _ := strconv.ParseUint(a, 2, 64)
	bi, _ := strconv.ParseUint(b, 2, 64)
	ci := ai + bi
	//cb, _ := strconv.ParseUint(strconv.Itoa(int(ci)), 2, 64)
	cs := strconv.FormatUint(ci, 2)
	return cs
}
