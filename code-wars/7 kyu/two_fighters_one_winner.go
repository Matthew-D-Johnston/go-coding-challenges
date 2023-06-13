package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	currentAttacker := firstAttacker

	for fighter1.Health > 0 && fighter2.Health > 0 {
		if fighter1.Name == currentAttacker {
			fighter2.Health -= fighter1.DamagePerAttack
			currentAttacker = fighter2.Name
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			currentAttacker = fighter1.Name
		}
	}

	if fighter1.Health <= 0 {
		return fighter2.Name
	}

	return fighter1.Name
}

func main() {
	fmt.Println(DeclareWinner(Fighter{Name: "Lew", Health: 10, DamagePerAttack: 1}, Fighter{Name: "Harry", Health: 10, DamagePerAttack: 4}, "Lew"))
}
