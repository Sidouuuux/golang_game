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
		p.position.y++
	case pos > 0:
		p.position.y--
	default:

	}
}
func (p *Player) askPlayerMove() {
	fmt.Print(p.name, " peut bouger ")

	if p.position.x > 0 {
		fmt.Print(" à gauche,")
	} else if p.position.x < 8 {
		fmt.Print(" à droite,")
	}
	if p.position.y > 0 {
		fmt.Print(" en haut,")
	} else if p.position.y < 4 {
		fmt.Print(" en bas,")
	}

	fmt.Print("\n\nPour bouger de haut en bas taper V\nPour bouger de droite a gauche taper H\n-> ")
	side := ""
	var choice int
	fmt.Scan(&side)

	switch {
	case side == "H":
		fmt.Print("Pour bouger a gauche taper -1, pour bouger a droite taper 1 -> ")
		fmt.Scan(&choice)
		p.movePlayerHorizontal(choice)
	case side == "V":
		fmt.Print("Pour monter taper 1, pour descendre taper -1 -> ")
		fmt.Scan(&choice)
		p.movePlayerVertical(choice)
	}
}
func printGrid(tab [5][9]string) {

	print("   ")
	for i := 65; i < (65 + 9); i++ {
		fmt.Printf("%c|", i)
	}
	print("\n")
	for i := 0; i < 5; i++ {
		fmt.Print("---------------------\n")
		fmt.Print(i+1, "|")
		for j := 0; j < 9; j++ {
			if tab[i][j] != "" {
				fmt.Print(tab[i][j], "|")
			} else {
				fmt.Print(" |")
			}
		}
		fmt.Print("\n")
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("*************************\n*Bienvenue dans le jeu !*\n*************************")

	var name1, name2 string
	fmt.Print("Entrez le nom du premier joueur et du deuxième joueur -> ")
	fmt.Scan(&name1)
	fmt.Print("Entrez le nom du deuxième joueur -> ")
	fmt.Scan(&name2)

	p1 := Player{name: name1, power: 100, position: Position{0, 0}}
	p2 := Player{name: name2, power: 100, position: Position{4, 8}}

	grid := [5][9]string{}
	tour := 1
	for p1.power > 0 && p2.power > 0 {
		grid[p1.position.x][p1.position.y] = "X"
		grid[p2.position.x][p2.position.y] = "O"
		printGrid(grid)
		fmt.Println("\nTour ", tour)
		fmt.Println("Vie ", p1.name, " = ", p1.power)
		fmt.Println("Vie ", p2.name, " = ", p2.power)
		switch {
		case tour%2 != 0:
			p1.askPlayerMove()
		default:
			p2.askPlayerMove()
			tour++

			grid[p1.position.x][p1.position.y] = ""
			grid[p2.position.x][p2.position.y] = ""
		}

	}
	fmt.Println("************************")
	fmt.Println("Vie ", p1.name, " = ", p1.power)
	fmt.Println("Vie ", p2.name, " = ", p2.power)
	if p1.power <= 0 {
		fmt.Println("Le joueur ", p2.name, " a gagné")
	} else {
		fmt.Println("Le joueur ", p1.name, " a gagné")
	}
}
