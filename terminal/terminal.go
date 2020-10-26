package terminal

import (
	"github.com/fatih/color"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// LineSize returns the terminal characters per line limit
func LineSize() int {
	var width int = 60

	x, err := terminal.Width()
	if err != nil {
		color.HiYellow("Couldn't get the terminal width")
	} else {
		width = int(x)
	}

	return width
}
