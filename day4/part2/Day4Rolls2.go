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

func findRolls(initialGrid []string) int {
	rows := len(initialGrid)
	if rows == 0 {
		return 0
	}
	cols := len(initialGrid[0])

	grid := make([][]byte, rows)
	for i := range grid {
		grid[i] = []byte(initialGrid[i])
	}

	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	countAdjacent := func(r, c int) int {
		count := 0
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
				count++
			}
		}
		return count
	}

	totalAccessible := 0
	changed := true

	for changed {
		changed = false
		var toRemove [][2]int

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == '@' && countAdjacent(r, c) < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		for _, pos := range toRemove {
			r, c := pos[0], pos[1]
			grid[r][c] = '.'
			totalAccessible++
			changed = true
		}
	}

	return totalAccessible
}
