package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsValidPassword(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc:        "[1,2,3] has a sum of 6.",
			input:       []string{`[1,2,3]`},
			expectedVal: 6,
		},
		{
			desc:        "{\"a\":2,\"b\":4} has a sum of 6.",
			input:       []string{`{"a":2,"b":4}`},
			expectedVal: 6,
		},
		{
			desc:        "[[[3]]] has a sum of 3.",
			input:       []string{`[[[3]]]`},
			expectedVal: 3,
		},
		{
			desc:        "{\"a\":{\"b\":4},\"c\":-1} has a sum of 3.",
			input:       []string{`{"a":{"b":4},"c":-1}`},
			expectedVal: 3,
		},
		{
			desc:        "{\"a\":[-1,1]} has a sum of 0.",
			input:       []string{`{"a":[-1,1]}`},
			expectedVal: 0,
		},
		{
			desc:        "[-1,{\"a\":1}] has a sum of 0",
			input:       []string{`[-1,{"a":1}]`},
			expectedVal: 0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumNumbers(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestIsValidPassword2(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc:        "[1,2,3] still has a sum of 6.",
			input:       []string{`[1,2,3]`},
			expectedVal: 6,
		},
		{
			desc:        "[1,{\"c\":\"red\",\"b\":2},3] now has a sum of 4, because the middle object is ignored.",
			input:       []string{`[1,{"c":"red","b":2},3]`},
			expectedVal: 4,
		},
		{
			desc:        "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5} now has a sum of 0, because the entire structure is ignored.",
			input:       []string{`{"d":"red","e":[1,2,3,4],"f":5}`},
			expectedVal: 0,
		},
		{
			desc:        "[1,\"red\",5] has a sum of 6, because \"red\" in an array has no effect.",
			input:       []string{`[1,"red",5]`},
			expectedVal: 6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumNumbers2(tc.input[0])
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
