package main

import (
	"testing"
)

var testInputFile = "../../data/1_test.txt"

func TestA(t *testing.T) {
	want := 24000
	have := solveA(testInputFile)
	if have != want {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := 24000 + 11000 + 10000
	have := solveB(testInputFile)
	if want != have {
		t.Fatalf("want %v != have %v\n", want, have)
	}
}
