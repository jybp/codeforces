package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	n := 0
	fmt.Fscan(r, &n)
	shortcuts := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &shortcuts[i])
	}
	fmt.Printf("shortcuts: %v\n", shortcuts)
}
