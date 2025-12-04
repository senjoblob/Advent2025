package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	accessible := findRolls(grid)
	fmt.Println(accessible)
}

func findRolls(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	accessible := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}

			adjacentRolls := 0

			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				accessible++
			}
		}
	}

	return accessible
}
