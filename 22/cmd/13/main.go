package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var day = 13

type Signal []interface{}

func parseSignal(in string) (signal []interface{}) {
	err := json.Unmarshal([]byte(in), &signal)
	if err != nil {
		panic(err)
	}
	return signal
}

func parseInput(f string) (packets [][2]Signal) {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	for _, block := range strings.Split(strings.Trim(string(input), "\n"), "\n\n") {
		packet := [2]Signal{}
		for i, signalString := range strings.Split(block, "\n") {
			packet[i] = parseSignal(signalString)
		}
		packets = append(packets, packet)
	}
	return packets
}

func cmpLists(a, b []interface{}) *bool {
	longer := a
	if len(b) > len(a) {
		longer = b
	}
	for i := range longer {
		if i >= len(a) {
			ordered := true
			return &ordered
		} else if i >= len(b) {
			ordered := false
			return &ordered
		}
		switch l := a[i].(type) {
		case float64:
			switch r := b[i].(type) {
			case float64:
				if l < r {
					ordered := true
					return &ordered
				} else if l > r {
					ordered := false
					return &ordered
				}
			case []interface{}:
				ordered := cmpLists([]interface{}{l}, r)
				if ordered != nil {
					return ordered
				}
			case interface{}:
				fmt.Println("whoopsie")
			}
		case []interface{}:
			switch r := b[i].(type) {
			case []interface{}:
				ordered := cmpLists(l, r)
				if ordered != nil {
					return ordered
				}
			case float64:
				ordered := cmpLists(l, []interface{}{r})
				if ordered != nil {
					return ordered
				}
			case interface{}:
				fmt.Println("whoopsie")
			}
		case interface{}:
			panic("whoopsie")
		}
	}
	return nil
}

func solveA(f string) (sum int) {
	packets := parseInput(f)
	for i, packet := range packets {
		signalsOrdered := cmpLists(packet[0], packet[1])
		if *signalsOrdered {
			sum += i + 1
		}
	}
	return sum
}

func solveB(f string) int {
	packets := parseInput(f)
	dividers := [2]Signal{
		{[]interface{}{float64(2)}},
		{[]interface{}{float64(6)}},
	}
	packets = append(packets, dividers)
	var packetsFlat [][]interface{}
	for _, packet := range packets {
		for _, sig := range packet {
			packetsFlat = append(packetsFlat, sig)
		}
	}

	sort.Slice(packetsFlat, func(i, j int) bool {
		return *cmpLists(packetsFlat[i], packetsFlat[j])
	})

	res := 1
	for i, sig := range packetsFlat {
		if SigEqual(sig, dividers[0]) || SigEqual(sig, dividers[1]) {
			res *= i + 1
		}
	}
	return res
}

func SigEqual(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		switch l := a[i].(type) {
		case []interface{}:
			switch r := b[i].(type) {
			case []interface{}:
				eq := SigEqual(l, r)
				if !eq {
					return false
				}
			case float64:
				return false
			}
		case float64:
			switch r := b[i].(type) {
			case []interface{}:
				return false
			case float64:
				if l != r {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
