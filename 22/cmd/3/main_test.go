package main

import "testing"

var testInputFile = "../../data/3_test.txt"

func TestA(t *testing.T) {
	want := 157
	have := solveA(testInputFile)
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := 70
	have := solveB(testInputFile)
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}
