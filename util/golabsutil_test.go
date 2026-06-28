package util

import (
	"strings"
	"testing"
)

func TestBanner(t *testing.T) {
	tests := []struct {
		name  string
		title string
		// wantMid is the middle line we expect, verifying padding + trimming.
		wantMid string
	}{
		{name: "simple", title: "Maps", wantMid: "│ Maps │"},
		{name: "trims surrounding space", title: "  Pointers  ", wantMid: "│ Pointers │"},
		{name: "empty", title: "", wantMid: "│  │"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lines := strings.Split(Banner(tc.title), "\n")
			if len(lines) != 3 {
				t.Fatalf("Banner(%q): got %d lines, want 3", tc.title, len(lines))
			}
			if lines[1] != tc.wantMid {
				t.Errorf("Banner(%q) middle line = %q, want %q", tc.title, lines[1], tc.wantMid)
			}
			// Top and bottom borders must match the middle line's width so the
			// box is rectangular regardless of title length.
			wantWidth := len([]rune(lines[1]))
			for _, i := range []int{0, 2} {
				if w := len([]rune(lines[i])); w != wantWidth {
					t.Errorf("Banner(%q) line %d width = %d, want %d", tc.title, i, w, wantWidth)
				}
			}
		})
	}
}
