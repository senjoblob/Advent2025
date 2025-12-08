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

	rows := len(grid)
	cols := len(grid[0])

	startCol := -1
	for c := 0; c < cols; c++ {
		if grid[0][c] == 'S' {
			startCol = c
			break
		}
	}

	active := make(map[int]bool)
	active[startCol] = true

	splits := 0

	for r := 0; r < rows; r++ {
		newActive := make(map[int]bool)

		for col := range active {
			if r+1 >= rows {
				continue
			}
			if grid[r+1][col] == '^' {
				splits++
				if col-1 >= 0 {
					newActive[col-1] = true
				}
				if col+1 < cols {
					newActive[col+1] = true
				}
			} else {
				newActive[col] = true
			}
		}
		active = newActive
	}

	fmt.Println(splits)
}
