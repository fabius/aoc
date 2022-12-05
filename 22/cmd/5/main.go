package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 5

type Command struct {
	amount int
	from   int
	to     int
}

func parseInput(f string) ([]string, []Command) {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	blocks := strings.Split(string(input), "\n\n")
	stacks := make([]string, 10)
	for _, l := range strings.Split(blocks[0], "\n") {
		if l == "" {
			continue
		}
		for i, j := 1, 0; i < len(l); i, j = i+4, j+1 {
			s := string(l[i])
			if s != "" && s != " " {
				stacks[j] = fmt.Sprintf("%s%s", s, stacks[j])
			}
		}
	}
	cmds := []Command{}
	for _, l := range strings.Split(blocks[1], "\n") {
		if l == "" {
			continue
		}
		cmd := strings.Split(l, " ")
		amount, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalln(err)
		}
		from, err := strconv.Atoi(cmd[3])
		if err != nil {
			log.Fatalln(err)
		}
		to, err := strconv.Atoi(cmd[5])
		if err != nil {
			log.Fatalln(err)
		}
		cmds = append(cmds, Command{
			amount: amount,
			from:   from,
			to:     to,
		})
	}
	return stacks, cmds
}

func solveA(f string) string {
	stacks, cmds := parseInput(f)
	for _, cmd := range cmds {
		for i := 1; i <= cmd.amount; i++ {
			moved := stacks[cmd.from-1][len(stacks[cmd.from-1])-1]
			stacks[cmd.from-1] = stacks[cmd.from-1][:len(stacks[cmd.from-1])-1]
			stacks[cmd.to-1] = fmt.Sprintf("%s%c", stacks[cmd.to-1], moved)
		}
	}
	res := ""
	for _, s := range stacks {
		if s == "" {
			continue
		}
		res += string(s[len(s)-1])
	}
	return res
}

func solveB(f string) string {
	stacks, cmds := parseInput(f)
	for _, cmd := range cmds {
		moved := string(stacks[cmd.from-1][len(stacks[cmd.from-1])-cmd.amount:])
		stacks[cmd.from-1] = stacks[cmd.from-1][:len(stacks[cmd.from-1])-cmd.amount]
		stacks[cmd.to-1] = fmt.Sprintf("%s%s", stacks[cmd.to-1], moved)
	}
	res := ""
	for _, s := range stacks {
		if s == "" {
			continue
		}
		res += string(s[len(s)-1])
	}
	return res
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %s\n", solveA(in))
	fmt.Printf("B: %s\n", solveB(in))
}
