package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var day = 12

type point struct {
	x, y int
}

// returns grid, start, end
func parseInput(f string) ([][]int, point, point) {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	grid := make([][]int, len(lines))
	var start, end point
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			if c == 'S' {
				start.y = i
				start.x = j
				c = 'a'
			} else if c == 'E' {
				end.y = i
				end.x = j
				c = 'z'
			}
			grid[i][j] = int(c)
		}
	}
	return grid, start, end
}

func climb(grid [][]int, start point, end point) (map[point]int, map[point]point) {
	dist := map[point]int{}
	prev := map[point]point{}
	for y, row := range grid {
		for x := range row {
			p := point{x: x, y: y}
			dist[p] = 1_000_000
		}
	}
	dist[start] = 0

	visited := map[point]bool{}
	pq := []point{start}
	diffs := []point{{x: -1}, {x: 1}, {y: -1}, {y: 1}}
	for len(pq) > 0 {
		p := pq[0]
		pq = pq[1:]
		if visited[p] {
			continue
		}
		visited[p] = true

		for _, d := range diffs {
			n := point{x: p.x - d.x, y: p.y - d.y}
			if visited[n] ||
				n.x < 0 || n.x >= len(grid[0]) ||
				n.y < 0 || n.y >= len(grid) {
				continue
			}
			currentHeight := grid[p.y][p.x]
			nextHeight := grid[n.y][n.x]
			if currentHeight+1 < nextHeight {
				continue
			}

			dcurrent := dist[p]
			dnext := dist[n]
			if dcurrent+1 < dnext {
				dist[n] = dcurrent + 1
				prev[n] = p
				pq = append(pq, n)
			}
		}
	}

	return dist, prev
}

func solveA(f string) int {
	grid, start, end := parseInput(f)
	dist, _ := climb(grid, start, end)
	return dist[end]
}

func solveB(f string) int {
	return 0
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
