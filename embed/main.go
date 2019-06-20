package main

import (
	"fmt"
)

func main() {
	myShip := Ship{"rifter", 60, Turret{Damage: 5}}
	targetShip := Ship{"merlin", 50, Turret{1}}
	myShip.Fire(&targetShip)
	fmt.Printf("Health remaining: %d \n", targetShip.Health)
}

type Turret struct {
	Damage int
}

func (t *Turret) Fire(target *Ship) {
	target.Health -= t.Damage
}

type Ship struct {
	Name   string
	Health int
	Turret
}
