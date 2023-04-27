package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

type shortcut struct {
	from, to int
}

func run(r io.Reader, w io.Writer) {
	n := 0
	fmt.Fscan(r, &n)

	shortcuts := make([]shortcut, n)
	for i := 0; i < n; i++ {
		shortcuts[i].from = i + 1
		fmt.Fscan(r, &shortcuts[i].to)
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(w, "%d ", solve(i, shortcuts))
	}
	fmt.Printf("shortcuts: %v\n", shortcuts)
}

func solve(i int, shortcuts []shortcut) int {
	// Withupt using shortcuts.
	return int(math.Abs(float64(i) - 1))
}
