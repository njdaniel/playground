package main

import (
	"fmt"
	"sort"
)

//Interval is the start and end time
type Interval struct {
	Start int
	End   int
}

func main() {
	intervals := make([]Interval, 0)
	i1 := Interval{
		Start: 20,
		End:   45,
	}
	i2 := Interval{
		Start: 12,
		End:   12,
	}
	i3 := Interval{
		Start: 2,
		End:   50,
	}
	i4 := Interval{
		Start: 14,
		End:   20,
	}
	i5 := Interval{
		Start: 3,
		End:   5,
	}
	intervals = append(intervals, i1, i2, i3, i4, i5)
	fmt.Println(minMeetingRooms(intervals))
}

func minMeetingRooms(intervals []Interval) int {
	n := len(intervals)
	if n == 0 {
		return n
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start <= intervals[j].Start
	})

	rooms := []int{}
	rooms = append(rooms, intervals[0].End)

	for _, interval := range intervals[1:] {
		sort.Slice(rooms, func(i, j int) bool {
			return rooms[i] < rooms[j]
		})

		switch {
		case rooms[0] <= interval.Start:
			rooms[0] = interval.End
		default:
			rooms = append(rooms, interval.End)
		}
	}

	return len(rooms)
}
