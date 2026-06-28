// Package util holds small shared helpers used across the go-labs examples.
//
// The examples in this repository are deliberately self-contained, but a few
// presentational helpers recur often enough to live in one place. Keeping them
// here also gives the module a non-main package to exercise with `go test`.
package util

import (
	"fmt"
	"strings"
)

// Banner returns a boxed, multi-line title string. It performs no I/O, which
// keeps it trivially testable; call Section to print it.
//
//	┌────────────┐
//	│  My title  │
//	└────────────┘
func Banner(title string) string {
	title = strings.TrimSpace(title)
	inner := len(title) + 2 // one space of padding on each side
	top := "┌" + strings.Repeat("─", inner) + "┐"
	mid := "│ " + title + " │"
	bot := "└" + strings.Repeat("─", inner) + "┘"
	return strings.Join([]string{top, mid, bot}, "\n")
}

// Section prints a Banner for title to stdout, followed by a blank line. Handy
// for visually separating the output of the small example programs.
func Section(title string) {
	fmt.Println(Banner(title))
	fmt.Println()
}
