package main

import (
	"fmt"
)

func main() {
	intervals := make([]Interval, 0)
	i1 := Interval{
		Start: 0,
		End:   30,
	}
	i2 := Interval{
		Start: 5,
		End:   10,
	}
	i3 := Interval{
		Start: 15,
		End:   20,
	}
	intervals = append(intervals, i1, i2, i3)
	fmt.Println(minMeetingRooms(intervals))
}

//Interval is the start and end time
type Interval struct {
	Start int
	End   int
}

//find the min number of meeting rooms needed
func minMeetingRooms(intervals []Interval) int {
	min := 0
	rooms := make([][]int, 0)
	for _, i := range intervals {
		addRoom := true
		for _, r := range rooms {
			if i.Start > r[0] && i.End > r[len(r)-1] {
				r = append(r, i.Start, i.End)
				addRoom = false
				break
			} else if i.Start < r[0] && i.End < r[len(r)-1] {
				is := []int{i.Start, i.End}
				r = append(is, r...)
				addRoom = false
				break
			}
		}
		if addRoom {
			room := make([]int, 0)
			room = append(room, i.Start, i.End)
			rooms = append(rooms, room)
		}
		min = len(rooms)

	}
	return min
}
