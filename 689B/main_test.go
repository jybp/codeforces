package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input: `3
2 2 3`,
			output: `0 1 2`,
		},
		{
			input: `5
1 2 3 4 5`,
			output: `0 1 2 3 4`,
		},
		{
			input: `7
		4 4 4 4 7 7 7`,
			output: `0 1 2 1 2 3 3`,
		},
	}
	for i, tc := range testCases {
		r := strings.NewReader(tc.input)
		var b bytes.Buffer
		run(r, &b)

		expected := tc.output
		actual := b.String()
		if expected != actual {
			t.Fatalf("test case %d failed: actual:\n%s\nexpected:\n%s",
				i+1, actual, expected)
		}
	}
}
