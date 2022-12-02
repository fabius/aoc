package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInput(f string) []string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	matches := []string{}
	for scanner.Scan() {
		lineTxt := scanner.Text()
		matches = append(matches, lineTxt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return matches
}

// returns my score
func solveA(f string) int {
	// opponent/me
	// A/X == rock == 1
	// B/Y == paper == 2
	// C/Z == scissors == 3
	// +0 if lost
	// +3 if draw
	// +6 if won
	matches := parseInput(f)
	score := 0
	for _, m := range matches {
		switch string(m[2]) {
		case "X":
			score += 1
			if string(m[0]) == "A" {
				score += 3
			} else if string(m[0]) == "C" {
				score += 6
			}
		case "Y":
			score += 2
			if string(m[0]) == "B" {
				score += 3
			} else if string(m[0]) == "A" {
				score += 6
			}
		case "Z":
			score += 3
			if string(m[0]) == "C" {
				score += 3
			} else if string(m[0]) == "B" {
				score += 6
			}
		}
	}
	return score
}

func solveB(f string) int {
	// 2nd col
	// X means you need to lose
	// Y means you need to draw
	// Z means you need to win
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	decider := map[string]int{
		"A X": scores["Z"],
		"B X": scores["X"],
		"C X": scores["Y"],
		"A Y": scores["X"] + 3,
		"B Y": scores["Y"] + 3,
		"C Y": scores["Z"] + 3,
		"A Z": scores["Y"] + 6,
		"B Z": scores["Z"] + 6,
		"C Z": scores["X"] + 6,
	}

	matches := parseInput(f)
	score := 0
	for _, m := range matches {
		score += decider[m]
	}
	return score
}

func main() {
	fmt.Printf("%d\n", solveA("22/data/2.txt"))
	fmt.Printf("%d\n", solveB("22/data/2.txt"))
}
