package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMaximumHappiness(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc: "After trying every other seating arrangement in this hypothetical scenario, you find that this one is the most optimal, with a total change in happiness of 330.",
			input: []string{
				"Alice would gain 54 happiness units by sitting next to Bob.",
				"Alice would lose 79 happiness units by sitting next to Carol.",
				"Alice would lose 2 happiness units by sitting next to David.",
				"Bob would gain 83 happiness units by sitting next to Alice.",
				"Bob would lose 7 happiness units by sitting next to Carol.",
				"Bob would lose 63 happiness units by sitting next to David.",
				"Carol would lose 62 happiness units by sitting next to Alice.",
				"Carol would gain 60 happiness units by sitting next to Bob.",
				"Carol would gain 55 happiness units by sitting next to David.",
				"David would gain 46 happiness units by sitting next to Alice.",
				"David would lose 7 happiness units by sitting next to Bob.",
				"David would gain 41 happiness units by sitting next to Carol.",
			},
			expectedVal: 330,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMaximumHappiness(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
