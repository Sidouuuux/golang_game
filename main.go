package main

import (
	"fmt"
	"math/rand"
	"strings"
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
func (p *Player) verify_move(_key string, _possible_move string) bool {
	if !(strings.Contains(_possible_move, _key)) {
		return false
	}
	p.movePlayer(_key)
	return true
}

func (p *Player) movePlayer(_key string) {
	switch {
	case _key == "Z":
		p.position.y -= 1
	case _key == "S":
		p.position.y += 1
	case _key == "Q":
		p.position.x -= 1
	case _key == "D":
		p.position.x += 1
	}
	fmt.Printf("\n---Position de %s : %i, %i ---", p.name, p.position.x, p.position.y)
}

func (p *Player) askPlayerMove() {
	fmt.Print(p.name, " peut bouger")
	key := ""
	possible_move := ""
	if p.position.x > 0 {
		fmt.Print(" à gauche")
		possible_move += "Q"
	} else if p.position.x < 8 {
		fmt.Print(" à droite ")
		possible_move += "D"
	}
	if p.position.y > 0 {
		fmt.Print(" en haut")
		possible_move += "Z"
	} else if p.position.y < 4 {
		fmt.Print(" en bas")
		possible_move += "S"
	}

	fmt.Print("\n\nPour bouger appuyer sur Z, Q, S, D -> ")
	fmt.Scan(&key)

	can_move := false
	can_move = p.verify_move(key, possible_move)
	if can_move != true {
		fmt.Println("can't move")
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

	tour := 1
	for {
		grid := [5][9]string{}
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
		tour++
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
