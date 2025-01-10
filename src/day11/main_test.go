package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsValidPassword(t *testing.T) {
	tcs := []struct {
		desc        string
		input       string
		expectedVal bool
	}{
		{
			desc:        "hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l)",
			input:       "hijklmmn",
			expectedVal: false,
		},
		{
			desc:        "abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.",
			input:       "abbceffg",
			expectedVal: false,
		},
		{
			desc:        "abbcegjk fails the third requirement, because it only has one double letter (bb)",
			input:       "abbcegjk",
			expectedVal: false,
		},
		{
			desc:        "abcdffaa is OK",
			input:       "abcdffaa",
			expectedVal: true,
		},
		{
			desc:        "ghjaabcc is OK",
			input:       "ghjaabcc",
			expectedVal: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := isValidPassword(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNextPassword(t *testing.T) {
	tcs := []struct {
		desc        string
		input       string
		expectedVal string
	}{
		{
			desc:        "The next password after abcdefgh is abcdffaa.",
			input:       "abcdefgh",
			expectedVal: "abcdffaa",
		},
		{
			desc:        "The next password after ghijklmn is ghjaabcc, because you eventually skip all the passwords that start with ghi..., since i is not allowed.",
			input:       "ghijklmn",
			expectedVal: "ghjaabcc",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNextPassword(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
