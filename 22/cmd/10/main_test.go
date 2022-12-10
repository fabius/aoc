package main

import (
	"fmt"
	"testing"
)

var testInputFile = fmt.Sprintf("../../data/%d_test.txt", day)

func TestA(t *testing.T) {
	want := 13140
	have := solveA(testInputFile)
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}
	have := solveB(testInputFile)
	for i := range want {
		if want[i] != have[i] {
			t.Fatalf("want != have \n%s\n%s\n", want[i], have[i])
		}
	}
}
