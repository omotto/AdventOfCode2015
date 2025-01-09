package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetLowestPositiveNumber(t *testing.T) {
	tcs := []struct {
		desc        string
		input       string
		numZeros    int
		expectedVal int
	}{
		{
			desc:        "If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.",
			input:       "abcdef",
			numZeros:    5,
			expectedVal: 609043,
		},
		{
			desc:        "If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....",
			input:       "pqrstuv",
			numZeros:    5,
			expectedVal: 1048970,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getLowestPositiveNumber(tc.input, tc.numZeros)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
