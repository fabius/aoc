package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 14

func parseInput(f string) (map[[2]int]bool, int) {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	grid := map[[2]int]bool{}
	highestY := 0
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, ">", "")
		tuples := strings.Split(line, "-")
		for i := range tuples {
			if i == len(tuples)-1 {
				break
			}

			tl := strings.Split(tuples[i], ",")
			xl, _ := strconv.Atoi(tl[0])
			yl, _ := strconv.Atoi(tl[1])

			tr := strings.Split(tuples[i+1], ",")
			xr, _ := strconv.Atoi(tr[0])
			yr, _ := strconv.Atoi(tr[1])

			for xi, yi := xl, yl; ; {
				p := [2]int{xi, yi}
				grid[p] = true
				if xr > xl {
					xi++
				} else if xr < xl {
					xi--
				}

				if yr > yl {
					yi++
				} else if yr < yl {
					yi--
				}

				if yi > highestY {
					highestY = yi
				}

				if xi == xr && yi == yr {
					p := [2]int{xi, yi}
					grid[p] = true
					break
				}
			}
		}
	}
	return grid, highestY + 2
}

func solve(grid map[[2]int]bool, floor *int) (sum int) {
outer:
	for ; ; sum++ {
		for x, y := 500, 0; ; y++ {
			if floor == nil && y > 500 {
				// fall into the abyss
				break outer
			}

			if floor != nil && *floor == y+1 {
				// come to rest
				grid[[2]int{x, y}] = true
				break
			}

			if !grid[[2]int{x, y + 1}] {
				// fall straight down
				continue
			}

			downLeft := grid[[2]int{x - 1, y + 1}]
			if !downLeft {
				// fall down-left
				x--
				continue
			}

			downRight := grid[[2]int{x + 1, y + 1}]
			if !downRight {
				// fall down-right
				x++
				continue
			}

			if downLeft && downRight {
				// come to rest
				grid[[2]int{x, y}] = true
				if x == 500 && y == 0 {
					// clog up source
					sum++
					break outer
				}
				break
			}
		}
	}
	return sum
}

func solveA(f string) (sum int) {
	grid, _ := parseInput(f)
	return solve(grid, nil)
}

func solveB(f string) (sum int) {
	grid, floor := parseInput(f)
	return solve(grid, &floor)
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
