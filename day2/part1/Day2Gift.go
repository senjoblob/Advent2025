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
	scanner.Scan()
	input := scanner.Text()

	input = strings.TrimSuffix(input, ",")

	ranges := strings.Split(input, ",")

	total := 0

	for _, ids := range ranges {
		parts := strings.Split(ids, "-")
		if len(parts) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		for id := start; id <= end; id++ {
			if isInvalidID(id) {
				total += id
			}
		}
	}

	fmt.Println(total)
}

func isInvalidID(id int) bool {
	number := strconv.Itoa(id)

	if len(number)%2 != 0 {
		return false
	}

	half := len(number) / 2
	return number[:half] == number[half:]
}
