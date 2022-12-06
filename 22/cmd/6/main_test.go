package main

import (
	"fmt"
	"testing"
)

var testInputFile = fmt.Sprintf("../../data/%d_test.txt", day)

func TestA(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		have int
	}{
		{
			desc: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 7,
			have: solveA("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
		},
		{
			desc: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 5,
			have: solveA("bvwbjplbgvbhsrlpgdmjqwftvncz"),
		},
		{
			desc: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 6,
			have: solveA("nppdvjthqldpwncqszvftbrmjlhg"),
		},
		{
			desc: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 10,
			have: solveA("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
		},
		{
			desc: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 11,
			have: solveA("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
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

func TestB(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		have int
	}{
		{
			desc: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 19,
			have: solveB("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
		},
		{
			desc: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 23,
			have: solveB("bvwbjplbgvbhsrlpgdmjqwftvncz"),
		},
		{
			desc: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 23,
			have: solveB("nppdvjthqldpwncqszvftbrmjlhg"),
		},
		{
			desc: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 29,
			have: solveB("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
		},
		{
			desc: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 26,
			have: solveB("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
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
