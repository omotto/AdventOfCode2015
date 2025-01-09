package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumHouses(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "> delivers presents to 2 houses: one at the starting location, and one to the east.",
			inputList: []string{
				">",
			},
			expectedVal: 2,
		},
		{
			desc: "^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.",
			inputList: []string{
				"^>v<",
			},
			expectedVal: 4,
		},
		{
			desc: "^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.",
			inputList: []string{
				"^v^v^v^v^v",
			},
			expectedVal: 2,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumHouses(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumHouses2(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.",
			inputList: []string{
				"^v",
			},
			expectedVal: 3,
		},
		{
			desc: "^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.",
			inputList: []string{
				"^>v<",
			},
			expectedVal: 3,
		},
		{
			desc: "^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.",
			inputList: []string{
				"^v^v^v^v^v",
			},
			expectedVal: 11,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumHouses2(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
