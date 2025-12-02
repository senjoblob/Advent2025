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

	var total int64 = 0

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
			if isInvalidID(strconv.Itoa(id)) {
				total += int64(id)
			}
		}
	}

	fmt.Println(total)
}

func isInvalidID(id string) bool {
	length := len(id)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}

		pattern := id[:patternLen]
		valid := true

		for i := patternLen; i < length; i += patternLen {
			if id[i:i+patternLen] != pattern {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}
