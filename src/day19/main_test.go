package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumMolecules(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc: "So, in the example above, there are 4 distinct molecules (not five, because HOOH appears twice) after one replacement from HOH.",
			input: []string{
				"H => HO",
				"H => OH",
				"O => HH",
				"",
				"HOH",
			},
			expectedVal: 4,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumMolecules(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumSteps(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc: "So, you could make HOH after 3 steps.",
			input: []string{
				"e => H",
				"e => O",
				"H => HO",
				"H => OH",
				"O => HH",
				"",
				"HOH",
			},
			expectedVal: 3,
		},
		{
			desc: "Santa's favorite molecule, HOHOHO, can be made in 6 steps.",
			input: []string{
				"e => H",
				"e => O",
				"H => HO",
				"H => OH",
				"O => HH",
				"",
				"HOHOHO",
			},
			expectedVal: 6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumSteps(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
