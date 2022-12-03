package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func parseInput(f string) []string {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(input), "\n")
}

func runeInt(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 38
	} else {
		return int(r) - 96
	}
}

func solveA(f string) int {
	sum := 0
	for _, line := range parseInput(f) {
		first := line[:len(line)/2]
		second := line[len(line)/2:]
		m := map[string]bool{}
		for _, c := range first {
			m[string(c)] = true
		}
		for _, c := range second {
			if m[string(c)] == true {
				sum += runeInt(c)
				break
			}
		}
	}
	return sum
}

func runeIsUsed(r rune, rs []rune) bool {
	alreadyUsed := false
	for _, c := range rs {
		if r == c {
			alreadyUsed = true
		}
	}
	return alreadyUsed
}

func evalGroup(groupLines []string, groupId int) int {
	m := map[rune]int{}
	for _, gl := range groupLines {
		used := []rune{}
		for _, c := range gl {
			if !runeIsUsed(c, used) {
				m[c] += 1
				used = append(used, c)
			}
		}
	}
	for r, c := range m {
		if c == 3 {
			return runeInt(r)
		}
	}
	return 0
}

func solveB(f string) int {
	sum := 0
	groupId := 0
	groupLines := []string{}
	for i, line := range parseInput(f) {
		groupLines = append(groupLines, line)
		if (i+1)%3 == 0 {
			sum += evalGroup(groupLines, groupId)
			groupLines = []string{}
			groupId++
		}
	}
	return sum
}

func main() {
	fmt.Printf("%d\n", solveA("data/3.txt"))
	fmt.Printf("%d\n", solveB("data/3.txt"))
}
