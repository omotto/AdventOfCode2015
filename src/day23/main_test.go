package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetHouseNumber(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		expectedVal int
	}{
		{
			desc: "For example, this program sets b to 2, because the jio instruction causes it to skip the tpl instruction.",
			input: []string{
				"inc b",
				"jio b, +2",
				"tpl b",
				"inc b",
			},
			expectedVal: 2,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getRegisterB(tc.input)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
