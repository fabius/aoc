package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 10

type instruction struct {
	operation string
	amount    int
}

func parseInput(f string) []instruction {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	instructions := []instruction{}
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		l := strings.Split(line, " ")
		op := l[0]
		var amount int
		if op == "addx" {
			amount, _ = strconv.Atoi(l[1])
		}
		instructions = append(instructions, instruction{operation: op, amount: amount})
	}
	return instructions
}

func forCycle(cycle int, register int) int {
	for i := 0; i <= 5; i++ {
		if cycle == 20+i*40 {
			return cycle * register
		}
	}
	return 0
}

func solveA(f string) int {
	sum := 0
	register := 1
	cycle := 1
	for _, instr := range parseInput(f) {
		if instr.operation == "noop" {
			sum += forCycle(cycle, register)
			cycle++
		} else if instr.operation == "addx" {
			for i := 0; i < 2; i++ {
				sum += forCycle(cycle, register)
				cycle++
			}
			register += instr.amount
		}
	}
	return sum
}

func forSprite(cycle, register int) string {
	if cycle%40 == register-1 || cycle%40 == register || cycle%40 == register+1 {
		return "#"
	} else {
		return "."
	}
}

func solveB(f string) []string {
	crt := []string{""}
	crtIndex := 0

	register := 1
	cycle := 1
	for _, instr := range parseInput(f) {
		if instr.operation == "noop" {
			crt[crtIndex] += forSprite(cycle-1, register)
			cycle++
			if (cycle-1)%40 == 0 {
				crt = append(crt, "")
				crtIndex++
			}
		} else if instr.operation == "addx" {
			for i := 0; i < 2; i++ {
				crt[crtIndex] += forSprite(cycle-1, register)
				cycle++
				if (cycle-1)%40 == 0 {
					crt = append(crt, "")
					crtIndex++
				}
			}
			register += instr.amount
		}
	}
	return crt
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	for _, line := range solveB(in) {
		fmt.Printf("B: %s\n", line)
	}
}
