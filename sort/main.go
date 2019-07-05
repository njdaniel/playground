package main

import (
	"fmt"
	"sort"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.000Z"

func main() {
	var dates []time.Time
	datea, _ := time.Parse(timeFormat, "2015-09-01T16:10:00.000Z")

	dateb, _ := time.Parse(timeFormat, "2015-09-01T16:00:00.000Z")
	datec, _ := time.Parse(timeFormat, "2015-09-01T16:20:00.000Z")
	//fmt.Println(datea)

	dates = append(dates, datea, dateb, datec)
	fmt.Println(dates)
	sort.Slice(dates, func(i, j int) bool { return dates[i].Before(dates[j])})
	fmt.Println(dates)
}


