package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ranges [][2]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			ranges = append(ranges, [2]int{start, end})
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][2]int{}
	if len(ranges) > 0 {
		merged = append(merged, ranges[0])
		for i := 1; i < len(ranges); i++ {
			last := &merged[len(merged)-1]
			curr := ranges[i]

			if curr[0] <= last[1]+1 {
				if curr[1] > last[1] {
					last[1] = curr[1]
				}
			} else {
				merged = append(merged, curr)
			}
		}
	}

	freshIDs := 0
	for _, r := range merged {
		freshIDs += r[1] - r[0] + 1
	}

	fmt.Println(freshIDs)
}
