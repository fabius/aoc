package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 8

func parseInput(f string) [][]int {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			height, _ := strconv.Atoi(string(c))
			grid[i][j] = height
		}
	}
	return grid
}

func solveA(f string) int {
	grid := parseInput(f)
	sum := 0
	for y, row := range grid {
		for x, tree := range row {
			if y == 0 || y == len(grid)-1 || x == 0 || x == len(row)-1 {
				sum++
				continue
			}

			leftTallest := 0
			for k := 0; k < x; k++ {
				if row[k] > leftTallest {
					leftTallest = row[k]
				}
			}

			rightTallest := 0
			for k := len(row) - 1; k > x; k-- {
				if row[k] > rightTallest {
					rightTallest = row[k]
				}
			}

			upTallest := 0
			for k := 0; k < y; k++ {
				if grid[k][x] > upTallest {
					upTallest = grid[k][x]
				}
			}

			downTallest := 0
			for k := len(grid) - 1; k > y; k-- {
				if grid[k][x] > downTallest {
					downTallest = grid[k][x]
				}
			}

			if tree > leftTallest || tree > rightTallest || tree > upTallest || tree > downTallest {
				sum++
			}
		}
	}
	return sum
}

func solveB(f string) int {
	grid := parseInput(f)
	score := 0
	for y, row := range grid {
		for x, tree := range row {
			left := x
			for k := x - 1; k >= 0; k-- {
				if row[k] >= tree {
					left = x - k
					break
				}
			}

			right := len(row) - 1 - x
			for k := x + 1; k < len(row); k++ {
				if row[k] >= tree {
					right = k - x
					break
				}
			}

			up := y
			for k := y - 1; k >= 0; k-- {
				if grid[k][x] >= tree {
					up = y - k
					break
				}
			}

			down := len(grid) - 1 - y
			for k := y + 1; k < len(grid); k++ {
				if grid[k][x] >= tree {
					down = k - y
					break
				}
			}

			treeScore := left * right * up * down
			if treeScore > score {
				score = treeScore
			}
		}
	}
	return score
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
