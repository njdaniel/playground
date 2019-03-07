package main

import (
	"fmt"
)

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
		for k, r := range rooms {
			if i.Start >= r[len(r)-1] && i.End > r[len(r)-1] {
				rooms[k] = append(r, i.Start, i.End)
				addRoom = false
				break
			} else if i.Start < r[0] && i.End <= r[0] {
				is := []int{i.Start, i.End}
				rooms[k] = append(is, r...)
				addRoom = false
				break
			} else if len(r) >= 4 {
				j := 1
				for ; j < len(r)/2; j++ {
					if i.Start >= r[1+(j-1)*2] && i.End <= r[2+(j-1)*2] {
						is := []int{i.Start, i.End}
						rooms[k] = append(r[:2+(j-1)*2], append(is, r[2+(j-1)*2:]...)...)
						addRoom = false
						break
					}
				}
				if !addRoom {
					break
				}
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
