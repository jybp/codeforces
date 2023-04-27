package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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
	fmt.Printf("shortcuts: %v\n", shortcuts)
	answers := []string{}
	for i := 1; i <= n; i++ {
		answers = append(answers, strconv.Itoa(solve(i, shortcuts)))
	}
	fmt.Fprintf(w, "%s", strings.Join(answers, " "))
}

func solve(i int, shortcuts []shortcut) int {
	// Withupt using shortcuts.
	return int(math.Abs(float64(i) - 1))
}
