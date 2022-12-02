package main

import "testing"

func TestA(t *testing.T) {
	want := 15
	have := solveA("../../data/2_test.txt")
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := 12
	have := solveB("../../data/2_test.txt")
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}
