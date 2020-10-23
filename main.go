package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Position struct {
	x, y int
}

type Player struct {
	name     string
	power    int
	position Position
}

func newPlayer(_name string, _power int, _attack int) *Player {
	p := Player{name: _name, power: _power, position: Position{0, 0}}
	return &p
}
func generateDamage() int {
	return rand.Intn(50-10+1) + 10
}
func (p *Player) attackPlayer(p2 *Player) {
	attack := generateDamage()
	p.power -= attack

	fmt.Println(p2.name, " attaque ", p.name, " de ", attack)
}
func (p *Player) movePlayerHorizontal(pos int) {
	switch {
	case pos < 0:
		p.position.x--
	case pos > 0:
		p.position.x++
	default:

	}
}
func (p *Player) movePlayerVertical(pos int) {
	switch {
	case pos < 0:
		p.position.y--
	case pos > 0:
		p.position.y++
	default:

	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("*************************\n*Bienvenue dans le jeu !*\n*************************")

	var name1, name2 string
	fmt.Print("Entrez le nom du premier joueur et du deuxième joueur -> ")
	fmt.Scanf("%s %s", &name1, &name2)

	p1 := Player{name: name1, power: 100}
	p2 := Player{name: name2, power: 100}
	tour := 1
	for p1.power > 0 && p2.power > 0 {
		fmt.Println("\nTour ", tour)
		fmt.Println("Vie ", p1.name, " = ", p1.power)
		fmt.Println("Vie ", p2.name, " = ", p2.power)
		switch {
		case tour%2 == 0:
			p1.attackPlayer(&p2)
		default:
			p2.attackPlayer(&p1)
		}
		// fmt.Println("Position de ", p1.position)
		// fmt.Print("mouve a gauche")
		// p1.movePlayerVertical(-1)
		// fmt.Println("Position de ", p1.position)
		// p1.movePlayerVertical(1)
		// fmt.Println("Position de ", p1.position)
		tour++
	}
	fmt.Println("************************")
	fmt.Println("Vie ", p1.name, " = ", p1.power)
	fmt.Println("Vie ", p2.name, " = ", p2.power)
	if p1.power <= 0 {
		fmt.Println("Le joueur ", p2.name, " a gangé")
	} else {
		fmt.Println("Le joueur ", p1.name, " a gangé")
	}
}
