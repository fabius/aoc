package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 14

func parseInput(f string) map[[2]int]bool {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	grid := map[[2]int]bool{}
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

				if xi == xr && yi == yr {
					p := [2]int{xi, yi}
					grid[p] = true
					break
				}
			}
		}
	}
	return grid
}

func solveA(f string) (sum int) {
	grid := parseInput(f)
	fmt.Printf("grid: %v\n", grid)
outer:
	for ; ; sum++ {
		fmt.Printf("sum: %v\n", sum)
		for x, y := 500, 0; ; y++ {
			// fmt.Printf("%v, %v\n", x, y)
			if y > 500 {
				// fall into the abyss
				break outer
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
				break
			}
		}
	}
	return sum
}

func solveB(f string) int {
	return 0
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
