package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMinimumDistance(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc: "The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is 605 in this example.",
			input: []string{
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			},
			expectedVal: 605,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMinimumDistance(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetMaximumDistance(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc: "For example, given the distances above, the longest route would be 982 via (for example) Dublin -> London -> Belfast.",
			input: []string{
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			},
			expectedVal: 982,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMaximumDistance(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
