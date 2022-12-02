package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// return the highest sum of calories
func solveA(f string) int {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	highest := 0
	current := 0
	for scanner.Scan() {
		lineTxt := scanner.Text()
		if lineTxt == "" {
			current = 0
			continue
		}
		line, err := strconv.Atoi(lineTxt)
		if err != nil {
			log.Fatalln(err)
		}
		current += line
		if current > highest {
			highest = current
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return highest
}

// return the highest sum of calories
func solveB(f string) int {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	all := make(map[int]int)
	current := 0
	index := 0
	for scanner.Scan() {
		lineTxt := scanner.Text()
		if lineTxt == "" {
			current = 0
			index++
			continue
		}
		line, err := strconv.Atoi(lineTxt)
		if err != nil {
			log.Fatalln(err)
		}
		current += line
		all[index] = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	allArr := []int{}
	for _, v := range all {
		allArr = append(allArr, v)
	}
	sort.Slice(allArr, func(i, j int) bool {
		return allArr[i] > allArr[j]
	})
	sum := 0
	for _, v := range allArr[:3] {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(solveA("22/data/1.txt"))
	fmt.Println(solveB("22/data/1.txt"))
}
