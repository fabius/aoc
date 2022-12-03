package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var day = 0

func readInput(f string) []string {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(input), "\n")
}

func solveA(f string) int {
	return 0
}

func solveB(f string) int {
	return 0
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
