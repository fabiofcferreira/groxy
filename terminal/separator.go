package terminal

import (
	"fmt"

	"github.com/fatih/color"
)

// LineSeparator prints a whole line of a given character
func LineSeparator(ch string, color *color.Color, width int) {
	for i := 1; i <= width; i++ {
		color.Printf(ch)
	}

	fmt.Printf("\n")
}
