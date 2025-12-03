package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		best := 0
		maxOnes := int(line[len(line)-1] - '0')

		for i := len(line) - 2; i >= 0; i-- {
			currentDigit := int(line[i] - '0')

			joltage := 10*currentDigit + maxOnes
			if joltage > best {
				best = joltage
			}

			if currentDigit > maxOnes {
				maxOnes = currentDigit
			}
		}

		total += best
	}

	fmt.Println(total)
}
