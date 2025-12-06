package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func main() {
	fmt.Println(countFresh(os.Stdin))
}

func countFresh(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	ranges := compileRanges(scanner)

	return countFreshIDs(scanner, ranges)
}

func compileRanges(scanner *bufio.Scanner) []Range {
	var ranges []Range

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if r, ok := readRange(line); ok {
			ranges = append(ranges, r)
		}
	}

	return mergeRanges(ranges)
}

func readRange(line string) (Range, bool) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return Range{}, false
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return Range{}, false
	}

	return Range{Start: start, End: end}, true
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		if r.Start <= last.End+1 {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

func countFreshIDs(scanner *bufio.Scanner, ranges []Range) int {
	freshCount := 0

	for scanner.Scan() {
		if id, err := strconv.Atoi(scanner.Text()); err == nil {
			if isFresh(id, ranges) {
				freshCount++
			}
		}
	}

	return freshCount
}

func isFresh(id int, ranges []Range) bool {
	i := sort.Search(len(ranges), func(i int) bool {
		return ranges[i].Start > id
	})

	if i > 0 && id <= ranges[i-1].End {
		return true
	}

	return false
}
