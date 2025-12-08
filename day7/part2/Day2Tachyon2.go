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

	count := make([][]int, rows)
	for i := range count {
		count[i] = make([]int, cols)
	}
	count[0][startCol] = 1

	for r := 0; r < rows; r++ {
		changed := true
		for changed {
			changed = false
			for c := 0; c < cols; c++ {
				if count[r][c] > 0 && grid[r][c] == '^' {
					amount := count[r][c]
					if c-1 >= 0 {
						count[r][c-1] += amount
						changed = true
					}
					if c+1 < cols {
						count[r][c+1] += amount
						changed = true
					}
					count[r][c] = 0
				}
			}
		}

		if r+1 < rows {
			for c := 0; c < cols; c++ {
				if count[r][c] > 0 {
					count[r+1][c] += count[r][c]
					count[r][c] = 0
				}
			}
		}
	}

	total := 0
	for c := 0; c < cols; c++ {
		total += count[rows-1][c]
	}

	fmt.Println(total)
}
