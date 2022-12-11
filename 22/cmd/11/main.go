package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var day = 11

type monkey struct {
	inspected int
	items     []uint64
	inspect   func(uint64) uint64
	test      func(uint64) uint64
}

type Monkeys []monkey

func parseInput(f string) ([]monkey, uint64) {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	monkeys := []monkey{}
	biggest := uint64(1)
	blocks := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	for _, block := range blocks {
		m := monkey{}
		var divisor, iftrue, iffalse uint64
		for _, line := range strings.Split(block, "\n") {
			l := strings.ReplaceAll(line, ",", "")
			l = strings.TrimSpace(l)
			s := strings.Split(l, " ")

			switch s[0] + s[1] {
			case "Startingitems:":
				items := []uint64{}
				for _, itemStr := range s[2:] {
					item, _ := strconv.Atoi(itemStr)
					items = append(items, uint64(item))
				}
				m.items = items

			case "Operation:new":
				m.inspect = func(old uint64) uint64 {
					first := old
					second := old
					if s[3] != "old" {
						num, _ := strconv.Atoi(s[3])
						first = uint64(num)
					}
					if s[5] != "old" {
						num, _ := strconv.Atoi(s[5])
						second = uint64(num)
					}
					if s[4] == "*" {
						return first * second
					} else {
						return first + second
					}
				}

			case "Test:divisible":
				d, _ := strconv.Atoi(s[3])
				divisor = uint64(d)
				biggest *= divisor

			case "Iftrue:":
				t, _ := strconv.Atoi(s[5])
				iftrue = uint64(t)

			case "Iffalse:":
				f, _ := strconv.Atoi(s[5])
				iffalse = uint64(f)
				m.test = func(num uint64) uint64 {
					if num%divisor == 0 {
						return iftrue
					} else {
						return iffalse
					}
				}
			}
		}
		monkeys = append(monkeys, m)
	}

	return monkeys, biggest
}

func monkeySort(monkeys []monkey) []monkey {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	return monkeys
}

func (monkeys *Monkeys) inspectAndThrow(mi int, item uint64, b *uint64) {
	(*monkeys)[mi].inspected++
	item = (*monkeys)[mi].inspect(item)
	if b == nil {
		// A
		item /= 3
	} else {
		// B
		item = item % *b
	}
	throwTo := (*monkeys)[mi].test(item)
	(*monkeys)[throwTo].items = append((*monkeys)[throwTo].items, item)
	if len((*monkeys)[mi].items) > 1 {
		(*monkeys)[mi].items = (*monkeys)[mi].items[1:]
	} else {
		(*monkeys)[mi].items = []uint64{}
	}
}

func solve(monkeys Monkeys, rounds int, biggest *uint64) int {
	for i := 0; i < rounds; i++ {
		for mi := range monkeys {
			for _, item := range monkeys[mi].items {
				monkeys.inspectAndThrow(mi, item, biggest)
			}
		}
	}
	monkeys = monkeySort(monkeys)
	return monkeys[0].inspected * monkeys[1].inspected
}

func solveA(f string) int {
	monkeys, _ := parseInput(f)
	return solve(monkeys, 20, nil)
}

func solveB(f string) int {
	monkeys, biggest := parseInput(f)
	return solve(monkeys, 10_000, &biggest)
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
