package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetHouseNumber(t *testing.T) {
	tcs := []struct {
		desc        string
		input       string
		expectedVal int
	}{
		{
			desc:        "The fourth house gets 70 presents, because it is visited by Elves 1, 2, and 4, for a total of 10 + 20 + 40 = 70 presents.",
			input:       "70",
			expectedVal: 4,
		},
		{
			desc:        "The eighth house gets 150 presents: 8*10 + 4 * 10 + 2 * 10 + 1 * 10",
			input:       "150",
			expectedVal: 8,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getHouseNumber(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
