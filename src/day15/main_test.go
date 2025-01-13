package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetBestScore(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		calories    int
		expectedVal int
	}{
		{
			desc: "Multiplying these together (68 * 80 * 152 * 76, ignoring calories for now) results in a total score of 62842880.",
			input: []string{
				"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
				"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			},
			calories:    -1,
			expectedVal: 62842880,
		},
		{
			desc: "For example, given the ingredients above, if you had instead selected 40 teaspoons of butterscotch and 60 teaspoons of cinnamon (which still adds to 100), the total calorie count would be 40*8 + 60*3 = 500. The total score would go down, though: only 57600000, the best you can do in such trying circumstances.",
			input: []string{
				"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
				"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			},
			calories:    500,
			expectedVal: 57600000,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getBestScore(tc.input, tc.calories)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
