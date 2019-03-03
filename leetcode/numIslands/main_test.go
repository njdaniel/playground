package main

import "testing"


func TestnumIslands(t *testing.T) {
	input := [][]byte{
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0},
	}
	output :=numIslands(input)
	if output != 1 {
		t.Errorf("Got: %d \t Wanted: 1\n", output)
	}
}


