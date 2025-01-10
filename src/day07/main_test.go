package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetRegister(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		register    string
		expectedVal int
	}{
		{
			desc: "i: 65079",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "i",
			expectedVal: 65079,
		},
		{
			desc: "h: 65412",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "h",
			expectedVal: 65412,
		},
		{
			desc: "g: 114",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "g",
			expectedVal: 114,
		},
		{
			desc: "f: 492",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "f",
			expectedVal: 492,
		},
		{
			desc: "e: 507",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "e",
			expectedVal: 507,
		},
		{
			desc: "d: 72",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "d",
			expectedVal: 72,
		},
		{
			desc: "x: 123",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "x",
			expectedVal: 123,
		},
		{
			desc: "y: 456",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			register:    "y",
			expectedVal: 456,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			visited := make(map[string]int)
			got := getRegister(tc.input, tc.register, visited)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
