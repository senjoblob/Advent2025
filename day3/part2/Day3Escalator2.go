package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := int64(0)
	banksize := 12

	for scanner.Scan() {
		line := scanner.Text()

		joltage := maxJoltage(line, banksize)

		if num, err := strconv.ParseInt(joltage, 10, 64); err == nil {
			total += num
		} else {
			fmt.Printf("Error parsing %s: %v\n", joltage, err)
		}
	}

	fmt.Println(total)
}

func maxJoltage(batteries string, banksize int) string {
	n := len(batteries)
	if banksize >= n {
		return batteries
	}

	bank := make([]byte, 0, banksize)

	for i := 0; i < n; i++ {
		for len(bank) > 0 &&
			len(bank)+(n-i) > banksize &&
			batteries[i] > bank[len(bank)-1] {
			bank = bank[:len(bank)-1]
		}

		if len(bank) < banksize {
			bank = append(bank, batteries[i])
		}
	}
	return string(bank)
}
