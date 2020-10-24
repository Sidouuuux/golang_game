package main

import (
	"fmt"
)

func printGrid(tab [9][9]int) {
	for i := 0; i <= 9; i++ {
		fmt.Print("---------------------\n|")
		for j := 0; j <= 9; j++ {
			fmt.Print(tab[i][j], "|")
		}
		fmt.Print("\n")
	}
}
func main() {
	grid := [9][9]int{}
	printGrid(grid)

}
