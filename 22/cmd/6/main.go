package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var day = 6

func parseInput(f string) string {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(string(input))
}

func unique(s []byte) bool {
	keys := make(map[byte]bool)
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
		} else {
			return false
		}
	}
	return true
}

func solve(in string, bufSize int) int {
	buf := []byte(in[:bufSize])
	for i := bufSize; i < len(in); i++ {
		buf[i%bufSize] = in[i]
		if unique(buf) {
			return i + 1
		}
	}
	return 0
}

func solveA(in string) int {
	return solve(in, 4)
}

func solveB(in string) int {
	return solve(in, 14)
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	parsed := parseInput(in)
	fmt.Printf("A: %d\n", solveA(parsed))
	fmt.Printf("B: %d\n", solveB(parsed))
}
