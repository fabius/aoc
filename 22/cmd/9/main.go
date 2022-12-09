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

func solveA(f string) int {
	s := coordinate{x: 0, y: 0}
	h := s
	t := s
	tailPositions := map[coordinate]bool{t: true}
	for _, m := range parseInput(f) {
		for i := 1; i <= m.amount; i++ {
			h.move(m)
			t.follow(h)
			tailPositions[t] = true
		}
	}
	return len(tailPositions)
}

func solveB(f string) int {
	s := coordinate{x: 0, y: 0}
	knots := []coordinate{}
	for i := 0; i <= 9; i++ {
		knots = append(knots, s)
	}
	tailPositions := map[coordinate]bool{knots[9]: true}
	for _, m := range parseInput(f) {
		for i := 1; i <= m.amount; i++ {
			knots[0].move(m)
			for j := 1; j <= 9; j++ {
				knots[j].follow(knots[j-1])
			}
			tailPositions[knots[9]] = true
		}
	}
	return len(tailPositions)
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
