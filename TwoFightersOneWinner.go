package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for {
		if firstAttacker == fighter1.Name {
			fighter2.Health -= fighter1.DamagePerAttack
			firstAttacker = fighter2.Name
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			firstAttacker = fighter1.Name
		}

		if fighter1.Health <= 0 {
			return fighter2.Name
		}

		if fighter2.Health <= 0 {
			return fighter1.Name
		}
	}
	// var winner string

	// if firstAttacker == fighter1.Name {
	// 	for len(winner) < 1 {
	// 		fighter2.Health -= fighter1.DamagePerAttack
	// 		if fighter1.Health <= 0 || fighter2.Health <= 0 {
	// 			if fighter2.Health <= 0 {
	// 				winner = fighter1.Name
	// 			} else {
	// 				winner = fighter2.Name
	// 			}
	// 			break
	// 		}
	// 		fighter1.Health -= fighter2.DamagePerAttack
	// 		if fighter1.Health <= 0 || fighter2.Health <= 0 {
	// 			if fighter2.Health <= 0 {
	// 				winner = fighter1.Name
	// 			} else {
	// 				winner = fighter2.Name
	// 			}
	// 			break
	// 		}
	// 	}
	// } else {
	// 	for len(winner) < 1 {
	// 		fighter1.Health -= fighter2.DamagePerAttack
	// 		if fighter1.Health <= 0 || fighter2.Health <= 0 {
	// 			if fighter2.Health <= 0 {
	// 				winner = fighter1.Name
	// 			} else {
	// 				winner = fighter2.Name
	// 			}
	// 			break
	// 		}
	// 		fighter2.Health -= fighter1.DamagePerAttack
	// 		if fighter1.Health <= 0 || fighter2.Health <= 0 {
	// 			if fighter2.Health <= 0 {
	// 				winner = fighter1.Name
	// 			} else {
	// 				winner = fighter2.Name
	// 			}
	// 			break
	// 		}
	// 	}
	// }
	// return winner
}
func main() {
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
}
