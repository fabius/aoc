package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 4

type bounds struct {
	start int
	end   int
}

func parseInput(f string) []map[int]bounds {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	parsed := []map[int]bounds{}
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		b := map[int]bounds{}
		for i, pair := range strings.Split(line, ",") {
			n := strings.Split(pair, "-")
			start, err := strconv.Atoi(n[0])
			if err != nil {
				log.Fatalln(err)
			}
			end, err := strconv.Atoi(n[1])
			if err != nil {
				log.Fatalln(err)
			}
			b[i] = bounds{start: start, end: end}
			if i == 1 {
				parsed = append(parsed, b)
			}
		}
	}
	return parsed
}

func solveA(f string) int {
	fullyContained := 0
	for _, b := range parseInput(f) {
		if (b[0].start <= b[1].start && b[0].end >= b[1].end) ||
			(b[1].start <= b[0].start && b[1].end >= b[0].end) {
			fullyContained++
		}
	}
	return fullyContained
}

func solveB(f string) int {
	overlap := 0
	for _, b := range parseInput(f) {
		smEnd := b[0].end
		lgStart := b[1].start
		if b[0].start > b[1].start {
			lgStart = b[0].start
			smEnd = b[1].end
		}
		if lgStart <= smEnd {
			overlap++
		}
	}
	return overlap
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
