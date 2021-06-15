package main

import "fmt"

func main() {
	fmt.Println("starting ewar")
	a := []string{"riddick", "bagger", "itsme"}
	t := []string{"bob", "fred", "alf"}
	n := 2
	fmt.Println(AssignEwar(a, t, n))
}

//AssignEwar assigns targets to allies
//  assume allies and targets are unique
func AssignEwar(allies, targets []string, numTargets int) map[string][]string {

	alliesAssignedFoes := make(map[string][]string, 0)
	j := 0
	//	p := &targets[j]
	for _, v := range allies {
		for i := 0; i < numTargets; i++ {
			alliesAssignedFoes[v] = append(alliesAssignedFoes[v], targets[j])
			j++
			if j >= len(targets) {
				j = 0
			}
		}
	}

	return alliesAssignedFoes
}
