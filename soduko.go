package soduko

import (
	"fmt"
)

const M = 9

func Puzzle(a [][]int) {
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			fmt.Print(a[i][j])
			if j+1 != M {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Solve(grid [][]int, row int, col int, num int) bool {
	for x := 0; x < 9; x++ {
		if grid[row][x] == num {
			return false
		}
	}
	for x := 0; x < 9; x++ {
		if grid[x][col] == num {
			return false
		}
	}
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

func GoodStart(grid [][]int) bool {
	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid); row++ {
			for row2 := row + 1; row2 < len(grid); row2++ {
				if grid[row][col] == grid[row2][col] {
					return false
				}
			}
		}
	}
	return true
}

func Sudoku(grid [][]int, row int, col int) bool {
	if row == M-1 && col == M {
		return true
	}
	if col == M {
		row += 1
		col = 0
	}
	if grid[row][col] > 0 {
		return Sudoku(grid, row, col+1)
	}
	for num := 1; num <= M; num++ {
		if Solve(grid, row, col, num) {
			grid[row][col] = num
			if Sudoku(grid, row, col+1) {
				return true
			}
		}
		grid[row][col] = 0
	}
	return false
}
