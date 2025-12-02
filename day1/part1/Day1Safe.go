package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	var err error
	if len(os.Args) < 2 {
		if hasStdin() {
			input, err = readStdin()
			if err != nil {
				fmt.Printf("Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Usage: go run Day1Safe.go input.txt")
			fmt.Println("Or:    go run Day1Safe.go < input.txt")
			fmt.Println("Or:    go run Day1Safe.go \"L68\nL30\nR48\n...\"")
			return
		}
	} else {
		arg := os.Args[1]

		if _, err := os.Stat(arg); err == nil {
			input, err = readFile(arg)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}
		} else {
			input = arg
		}
	}

	password := solve(input)

	fmt.Printf("The password is: %d\n", password)
}

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func readStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(lines, "\n"), nil
}

func hasStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func solve(input string) int {
	position := 50
	count := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		direction := line[0]
		distance := 0
		fmt.Sscanf(line[1:], "%d", &distance)

		if direction == 'L' {
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		} else if direction == 'R' {
			position = (position + distance) % 100
		}

		if position == 0 {
			count++
		}
	}
	return count
}
