package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMinimumDistance(t *testing.T) {
	tcs := []struct {
		desc        string
		input       string
		times       int
		expectedVal int
	}{
		{
			desc:        "1 becomes 11 (1 copy of digit 1) -> 11 becomes 21 (2 copies of digit 1) -> 21 becomes 1211 (one 2 followed by one 1) -> 1211 becomes 111221 (one 1, one 2, and two 1s) -> 111221 becomes 312211 (three 1s, two 2s, and one 1)",
			input:       "1",
			times:       5,
			expectedVal: 6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getLookAndSayLength(tc.input, tc.times)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
