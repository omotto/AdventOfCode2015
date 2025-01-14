package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumLightOn(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		steps       int
		expectedVal int
	}{
		{
			desc: "After 4 steps, this example has four lights on.",
			input: []string{
				".#.#.#",
				"...##.",
				"#....#",
				"..#...",
				"#.#..#",
				"####..",
			},
			steps:       4,
			expectedVal: 4,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumLightOn(tc.input, tc.steps)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumLightOn2(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		steps       int
		expectedVal int
	}{
		{
			desc: "After 5 steps, this example now has 17 lights on.",
			input: []string{
				"##.#.#",
				"...##.",
				"#....#",
				"..#...",
				"#.#..#",
				"####.#",
			},
			steps:       5,
			expectedVal: 17,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumLightOn2(tc.input, tc.steps)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
