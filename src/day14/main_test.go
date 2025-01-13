package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMaxDistanceAt(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		seconds     int
		expectedVal int
	}{
		{
			desc: "In this example, after the 1000th second, both reindeer are resting, and Comet is in the lead at 1120 km (poor Dancer has only gotten 1056 km by that point). So, in this situation, Comet would win (if the race ended at 1000 seconds).",
			input: []string{
				"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
				"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
			},
			seconds:     1000,
			expectedVal: 1120,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMaxDistanceAt(tc.input, tc.seconds)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetMaxPointsAt(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		seconds     int
		expectedVal int
	}{
		{
			desc: "After the 1000th second, Dancer has accumulated 689 points, while poor Comet, our old champion, only has 312. So, with the new scoring system, Dancer would win (if the race ended at 1000 seconds).",
			input: []string{
				"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
				"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
			},
			seconds:     1000,
			expectedVal: 689,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMaxPointsAt(tc.input, tc.seconds)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
