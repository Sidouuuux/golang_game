package main

import (
	"fmt"
	"strings"
)

type Position struct {
	x, y int
}

type Player struct {
	name     string
	power    int
	position Position
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
}

func (p *Player) askPlayerMove() {
	fmt.Print(p.name, " peut bouger")
	key := ""
	possible_move := ""
	if p.position.x >= 1 {
		fmt.Print(" à gauche")
		possible_move += "Q"
	}
	if p.position.x < 8 {
		fmt.Print(" à droite")
		possible_move += "D"
	}
	if p.position.y >= 1 {
		fmt.Print(" en haut")
		possible_move += "Z"
	}
	if p.position.y < 4 {
		fmt.Print(" en bas")
		possible_move += "S"
	}

	fmt.Print("\n\nPour bouger appuyer sur Z, Q, S, D -> ")
	fmt.Scan(&key)
	key = strings.ToUpper(key)

	can_move := false
	can_move = p.verify_move(key, possible_move)
	if can_move != true {
		fmt.Println("can't move")
	}
}

func main() {

	p1 := Player{name: "oui", power: 100, position: Position{8, 4}}

	for {
		grid := [5][9]string{}
		grid[p1.position.y][p1.position.x] = "X"
		fmt.Printf("\n---Position de %s : %d/%d ---\n", p1.name, p1.position.x, p1.position.y)
		printGrid(grid)
		p1.askPlayerMove()
	}
}
