package terminal

import (
	"log"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

//TerminalSize returns the terminal characters per line limit
func TerminalSize() int {
	// get terminal size
	var width int = 120

	x, err := terminal.Width()
	if err != nil {
		log.Println("Couldn't get the terminal width")
	}
	width = int(x)

	return width
}
