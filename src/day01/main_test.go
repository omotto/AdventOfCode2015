package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetFloor(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "(()) result in floor 0",
			inputList: []string{
				"(())",
			},
			expectedVal: 0,
		},
		{
			desc: "()() result in floor 0",
			inputList: []string{
				"()()",
			},
			expectedVal: 0,
		},
		{
			desc: "((( result in floor 3",
			inputList: []string{
				"(((",
			},
			expectedVal: 3,
		},
		{
			desc: "(()(()( result in floor 3",
			inputList: []string{
				"(()(()(",
			},
			expectedVal: 3,
		},
		{
			desc: "))((((( also results in floor 3",
			inputList: []string{
				"))(((((",
			},
			expectedVal: 3,
		},
		{
			desc: "()) result in floor -1 (the first basement level).",
			inputList: []string{
				"())",
			},
			expectedVal: -1,
		},
		{
			desc: "))( result in floor -1 (the first basement level).",
			inputList: []string{
				"))(",
			},
			expectedVal: -1,
		},
		{
			desc: "))) result in floor -3",
			inputList: []string{
				")))",
			},
			expectedVal: -3,
		},
		{
			desc: ")())()) result in floor -3",
			inputList: []string{
				")())())",
			},
			expectedVal: -3,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getFloor(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetPosition(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: ") causes him to enter the basement at character position 1.",
			inputList: []string{
				")",
			},
			expectedVal: 1,
		},
		{
			desc: "()()) causes him to enter the basement at character position 5.",
			inputList: []string{
				"()())",
			},
			expectedVal: 5,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getPosition(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
