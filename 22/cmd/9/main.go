package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var day = 9

type motion struct {
	direction string
	amount    int
}

func parseInput(f string) []motion {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	motions := []motion{}
	for _, l := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		li := strings.Split(l, " ")
		amnt, _ := strconv.Atoi(li[1])
		motions = append(motions, motion{direction: li[0], amount: amnt})
	}
	return motions
}

type coordinate struct{ x, y int }

func (c *coordinate) move(m motion) {
	if m.direction == "R" {
		c.x++
	} else if m.direction == "L" {
		c.x--
	} else if m.direction == "U" {
		c.y++
	} else if m.direction == "D" {
		c.y--
	}
}

func (c *coordinate) follow(h coordinate) {
	d := math.Sqrt(math.Pow(float64(h.x-c.x), 2) + math.Pow(float64(h.y-c.y), 2))
	if d >= 2 {
		if h.y > c.y {
			c.y++
		} else if h.y < c.y {
			c.y--
		}
		if h.x > c.x {
			c.x++
		} else if h.x < c.x {
			c.x--
		}
	}
}

func solve(motions []motion, knotCount int) int {
	knots := []coordinate{}
	for i := 0; i <= knotCount-1; i++ {
		knots = append(knots, coordinate{x: 0, y: 0})
	}
	tailPositions := map[coordinate]bool{knots[knotCount-1]: true}
	for _, m := range motions {
		for i := 1; i <= m.amount; i++ {
			knots[0].move(m)
			for j := 1; j <= knotCount-1; j++ {
				knots[j].follow(knots[j-1])
			}
			tailPositions[knots[knotCount-1]] = true
		}
	}
	return len(tailPositions)
}

func solveA(f string) int {
	return solve(parseInput(f), 2)
}

func solveB(f string) int {
	return solve(parseInput(f), 10)
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
