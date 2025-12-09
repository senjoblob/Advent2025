package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Light struct {
	X, Y, Z int
}

type Connection struct {
	From, To int
	SqDist   int
}

type CircuitFind struct {
	parent []int
	size   []int
}

func NewCircuitFind(n int) *CircuitFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &CircuitFind{parent, size}
}

func (cf *CircuitFind) Find(x int) int {
	if cf.parent[x] != x {
		cf.parent[x] = cf.Find(cf.parent[x])
	}
	return cf.parent[x]
}

func (cf *CircuitFind) Circuit(x, y int) bool {
	originX := cf.Find(x)
	originY := cf.Find(y)

	if originX == originY {
		return false
	}

	if cf.size[originX] < cf.size[originY] {
		originX, originY = originY, originX
	}

	cf.parent[originY] = originX
	cf.size[originX] += cf.size[originY]
	return true
}

func (cf *CircuitFind) GetSize(x int) int {
	origin := cf.Find(x)
	return cf.size[origin]
}

func sqDistance(l1, l2 Light) int {
	dx := l1.X - l2.X
	dy := l1.Y - l2.Y
	dz := l1.Z - l2.Z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lights []Light

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			continue
		}
		lights = append(lights, Light{x, y, z})
	}

	n := len(lights)

	var connections []Connection
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := sqDistance(lights[i], lights[j])
			connections = append(connections, Connection{i, j, dist})
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].SqDist < connections[j].SqDist
	})

	cf := NewCircuitFind(n)
	for i := 0; i < 1000 && i < len(connections); i++ {
		conn := connections[i]
		cf.Circuit(conn.From, conn.To)
	}

	circuitSizes := make(map[int]int)
	for i := 0; i < n; i++ {
		origin := cf.Find(i)
		circuitSizes[origin]++
	}

	var sizes []int
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		fmt.Println(sizes[i])
		result *= sizes[i]
	}

	fmt.Println(result)
}
