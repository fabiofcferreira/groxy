package terminal

import (
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// LineSize returns the terminal characters per line limit
func LineSize() int {
	var width int = 60

	x, err := terminal.Width()
	if err == nil {
		width = int(x)
	}

	return width
}
