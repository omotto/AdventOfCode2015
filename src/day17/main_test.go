package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumCombinations(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		liters      int
		expectedVal int
	}{
		{
			desc:        "For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there are four ways to do it.",
			input:       []string{"20", "15", "10", "5", "5"},
			liters:      25,
			expectedVal: 4,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumCombinations(tc.input, tc.liters)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumMinimumWays(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		liters      int
		expectedVal int
	}{
		{
			desc:        "In the example above, the minimum number of containers was two. There were three ways to use that many containers, and so the answer there would be 3.",
			input:       []string{"20", "15", "10", "5", "5"},
			liters:      25,
			expectedVal: 3,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumMinimumWays(tc.input, tc.liters)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
