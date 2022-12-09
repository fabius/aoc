package main

import (
	"fmt"
	"testing"
)

var testInputFile = fmt.Sprintf("../../data/%d_test.txt", day)
var testInputFileB2 = fmt.Sprintf("../../data/%d_test_b2.txt", day)

func TestA(t *testing.T) {
	want := 13
	have := solveA(testInputFile)
	if want != have {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		have int
	}{
		{
			desc: testInputFile,
			want: 1,
			have: solveB(testInputFile),
		},
		{
			desc: testInputFileB2,
			want: 36,
			have: solveB(testInputFileB2),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.want != tC.have {
				t.Fatalf("want %d != have %d\n", tC.want, tC.have)
			}
		})
	}
}
