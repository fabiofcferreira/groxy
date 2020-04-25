package terminal

import "github.com/fatih/color"

//YesNoColored prints Yes in green and No in red
func YesNoColored(test bool) {
	if test {
		color.HiGreen("Yes")
	} else {
		color.HiRed("No")
	}
}
