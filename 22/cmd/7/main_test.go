package main

import (
	"fmt"
	"testing"
)

var testInputFile = fmt.Sprintf("../../data/%d_test.txt", day)

func TestA(t *testing.T) {
	want := 95437
	have := solveA(testInputFile)
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := 24933642
	have := solveB(testInputFile)
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}