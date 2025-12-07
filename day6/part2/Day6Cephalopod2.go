package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	total := 0
	maxLen := 0

	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	col := maxLen - 1
	for col >= 0 {
		empty := true
		for row := 0; row < len(lines)-1; row++ {
			if col < len(lines[row]) && lines[row][col] != ' ' {
				empty = false
				break
			}
		}

		if empty {
			col--
			continue
		}

		var numbers []int
		var operation byte

		for col >= 0 {
			separator := true
			for row := 0; row < len(lines)-1; row++ {
				if col < len(lines[row]) && lines[row][col] != ' ' {
					separator = false
					break
				}
			}

			if separator {
				break
			}

			var digits []byte
			for row := 0; row < len(lines)-1; row++ {
				if col < len(lines[row]) {
					ch := lines[row][col]
					if ch >= '0' && ch <= '9' {
						digits = append(digits, ch)
					}
				}
			}

			if len(digits) > 0 {
				num := 0
				for _, digit := range digits {
					num = num*10 + int(digit-'0')
				}
				numbers = append([]int{num}, numbers...)
			}

			if col < len(lines[len(lines)-1]) {
				ch := lines[len(lines)-1][col]
				if ch == '+' || ch == '*' {
					operation = ch
				}
			}

			col--
		}

		if len(numbers) > 0 && operation != 0 {
			result := numbers[0]
			for i := 1; i < len(numbers); i++ {
				if operation == '+' {
					result += numbers[i]
				} else {
					result *= numbers[i]
				}
			}

			total += result
		}
	}

	fmt.Println(total)
}
