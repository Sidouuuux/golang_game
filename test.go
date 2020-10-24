package main

import "fmt"

func printGrid(tab [5][9]string) {

	print("   ")
	for i := 65; i < (65 + 9); i++ {
		fmt.Printf("%c|", i)
	}
	print("\n")
	for i := 0; i < 5; i++ {
		fmt.Print("----------------------\n")
		fmt.Print(i, "| ")
		for j := 0; j < 9; j++ {
			fmt.Print(tab[i][j], "|")
		}
		fmt.Print("\n")
	}
}
func main() {
	grid := [5][9]string{}

	// fmt.Println(grid)
	printGrid(grid)

}
