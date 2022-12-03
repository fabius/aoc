package main

import "testing"

func TestA(t *testing.T) {
	want := 157
	have := solveA("../../data/3_test.txt")
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := 70
	have := solveB("../../data/3_test.txt")
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}
