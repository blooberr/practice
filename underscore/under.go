package main

import (
	"fmt"
)

// SimplifyUnderscore takes an input, and condenses consecutive underscores
// into one underscore
func SimplifyUnderscore(input string) (output string) {
	if len(input) == 0 {
		return output
	}

	underCounter := 0

	for _, character := range input {
		if string(character) == "_" {
			underCounter = underCounter + 1
		} else {
			underCounter = 0
		}

		if !(underCounter > 1) {
			output = fmt.Sprintf("%s%s", output, string(character))
		}
	}
	return output
}
