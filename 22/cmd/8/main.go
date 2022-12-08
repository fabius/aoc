package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 8

func readInput(f string) [][]int {
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
	grid := readInput(f)
	sum := 0
	for i, row := range grid {
		for j, tree := range row {
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(row)-1 {
				sum++
				continue
			}

			leftTallest := 0
			for k := 0; k < j; k++ {
				if row[k] > leftTallest {
					leftTallest = row[k]
				}
			}

			rightTallest := 0
			for k := len(row) - 1; k > j; k-- {
				if row[k] > rightTallest {
					rightTallest = row[k]
				}
			}

			upTallest := 0
			for k := 0; k < i; k++ {
				if grid[k][j] > upTallest {
					upTallest = grid[k][j]
				}
			}

			downTallest := 0
			for k := len(grid) - 1; k > i; k-- {
				if grid[k][j] > downTallest {
					downTallest = grid[k][j]
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
	grid := readInput(f)
	score := 0
	for i, row := range grid {
		for j, tree := range row {
			left := 0
			for k := j - 1; k >= 0; k-- {
				if row[k] >= tree {
					left = j - k
					break
				}
			}
			if left == 0 {
				left = j
			}

			right := 0
			for k := j + 1; k < len(row); k++ {
				if row[k] >= tree {
					right = k - j
					break
				}
			}
			if right == 0 {
				right = len(row) - 1 - j
			}

			up := 0
			for k := i - 1; k >= 0; k-- {
				if grid[k][j] >= tree {
					up = i - k
					break
				}
			}
			if up == 0 {
				up = i
			}

			down := 0
			for k := i + 1; k < len(grid); k++ {
				if grid[k][j] >= tree {
					down = k - i
					break
				}
			}
			if down == 0 {
				down = len(grid) - 1 - i
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
