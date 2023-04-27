package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	r := strings.NewReader(`3
	2 2 3`)
	var b bytes.Buffer
	run(r, &b)

	expected := `0 1 2 `
	actual := b.String()
	if expected != actual {
		t.Fatalf("actual:\n%s\nexpected:\n%s", actual, expected)
	}
}
