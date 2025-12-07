package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers [][]int
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	operations := strings.Fields(lines[len(lines)-1])

	for _, line := range lines[:len(lines)-1] {
		parts := strings.Fields(line)
		if numbers == nil {
			numbers = make([][]int, len(parts))
		}

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers[i] = append(numbers[i], num)
		}
	}

	total := 0
	for i, nums := range numbers {
		if i >= len(operations) || len(nums) == 0 {
			continue
		}

		result := nums[0]
		for j := 1; j < len(nums); j++ {
			if operations[i] == "+" {
				result += nums[j]
			} else {
				result *= nums[j]
			}
		}
		total += result
	}

	fmt.Println(total)
}
