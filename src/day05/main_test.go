package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNiceStringNumber(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc:        "ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.",
			input:       []string{"ugknbfddgicrmopn"},
			expectedVal: 1,
		},
		{
			desc:        "aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.",
			input:       []string{"aaa"},
			expectedVal: 1,
		},
		{
			desc:        "jchzalrnumimnmhp is naughty because it has no double letter.",
			input:       []string{"jchzalrnumimnmhp"},
			expectedVal: 0,
		},
		{
			desc:        "haegwjzuvuyypxyu is naughty because it contains the string xy.",
			input:       []string{"haegwjzuvuyypxyu"},
			expectedVal: 0,
		},
		{
			desc:        "dvszwmarrgswjxmb is naughty because it contains only one vowel.",
			input:       []string{"dvszwmarrgswjxmb"},
			expectedVal: 0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNiceStringNumber(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNiceStringNumber2(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc:        "qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).",
			input:       []string{"qjhvhtzxzqqjkmpb"},
			expectedVal: 1,
		},
		{
			desc:        "xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.",
			input:       []string{"xxyxx"},
			expectedVal: 1,
		},
		{
			desc:        "uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.",
			input:       []string{"uurcxstgmygtbstg"},
			expectedVal: 0,
		},
		{
			desc:        "ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.",
			input:       []string{"ieodomkazucvgmuy"},
			expectedVal: 0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNiceStringNumber2(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
