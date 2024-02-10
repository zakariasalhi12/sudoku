package main

import (
	"fmt"
	"os"
	"soduko"
)

const M = 9

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	grid := make([][]int, M)

	for i := 0; i < M; i++ {
		grid[i] = make([]int, M)
		for j, char := range os.Args[i+1] {
			if char == '.' {
				grid[i][j] = 0
			} else {
				grid[i][j] = int(char - '0')
			}
		}
	}

	if soduko.Sudoku(grid, 0, 0) && soduko.GoodStart(grid) {
		soduko.Puzzle(grid)
	} else {
		fmt.Println("Error")
	}
}
