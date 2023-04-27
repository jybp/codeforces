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

type node struct {
	position int
	next     *node
	shortcut *node
}

func run(r io.Reader, w io.Writer) {
	n := 0
	fmt.Fscan(r, &n)

	shortcuts := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		var to int
		fmt.Fscan(r, &to)
		shortcuts[i] = to
	}
	fmt.Printf("shortcuts: %v\n", shortcuts)

	// Setup regular routes.
	nodes := map[int]*node{}
	root := &node{}
	current := root
	for j := 1; j <= n; j++ {
		current.position = j
		nodes[j] = current
		next := &node{}
		if j < n {
			current.next = next
			current = next
		}
	}

	// Fill all shorcuts for each node.
	for j := 1; j <= n; j++ {
		if to, ok := shortcuts[j]; ok {
			nodes[j].shortcut = nodes[to]
		}
	}

	answers := []string{}
	for i := 1; i <= n; i++ {
		answers = append(answers, strconv.Itoa(solve(i, root)))
	}
	fmt.Fprintf(w, "%s", strings.Join(answers, " "))
}

// We want to to go from 1 -> i with the minimum amount of energy.
func solve(i int, root *node) int {

	// Without using shortcuts.
	// return int(math.Abs(float64(i) - 1))

	// Find the cheapest path from the root to node i.
	cost := 0
	current := root

	for {
		// We've reached the target.
		if current.position == i {
			break
		}
		// We've reach the end, backtrack.
		if current.next == nil {
			cost += int(math.Abs(float64(current.position) - float64(i)))
			break
		}
		// No shorcut, proceed.
		if current.shortcut == nil {
			cost += 1
			current = current.next
			continue
		}

		distWithShortcut, distWithoutShortcut := math.MaxInt, math.MaxInt
		distWithShortcut = int(math.Abs(float64(current.shortcut.position) - float64(i)))
		distWithoutShortcut = int(math.Abs(float64(current.next.position) - float64(i)))
		if distWithShortcut < distWithoutShortcut {
			current = current.shortcut
		} else {
			current = current.next
		}
		cost += 1
	}

	return cost
}
