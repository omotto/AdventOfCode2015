package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMiQuantum(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc:        "In this situation, the quantum entanglement for the ideal configuration is therefore 99.",
			input:       []string{"11", "9", "10", "8", "2", "7", "5", "4", "3", "1"},
			expectedVal: 99,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMiQuantum(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetMiQuantum2(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc:        "Of these, there are three arrangements that put the minimum (two) number of packages in the first group: 11 4, 10 5, and 8 7. Of these, 11 4 has the lowest quantum entanglement, and so it is selected.",
			input:       []string{"11", "9", "10", "8", "2", "7", "5", "4", "3", "1"},
			expectedVal: 44,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMiQuantum2(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
