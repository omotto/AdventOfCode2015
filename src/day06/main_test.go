package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumLightsBrightness(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc:        "turn on 0,0 through 0,0 would increase the total brightness by 1.",
			input:       []string{"turn on 0,0 through 0,0"},
			expectedVal: 1,
		},
		{
			desc:        "toggle 0,0 through 999,999 would increase the total brightness by 2000000.",
			input:       []string{"toggle 0,0 through 999,999"},
			expectedVal: 2000000,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumLightsBrightness(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
