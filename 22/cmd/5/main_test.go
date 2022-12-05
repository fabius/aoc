package main

import (
	"fmt"
	"testing"
)

var testInputFile = fmt.Sprintf("../../data/%d_test.txt", day)

func TestA(t *testing.T) {
	want := "CMZ"
	have := solveA(testInputFile)
	if want != have {
		t.Fatalf("want %s != have %s\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := "MCD"
	have := solveB(testInputFile)
	if want != have {
		t.Fatalf("want %s != have %s\n", want, have)
	}
}
