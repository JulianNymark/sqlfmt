package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tokens := tokenizeInput()
	formatted := formatTokens(tokens)
	fmt.Print(formatted)
}

func formatTokens(tokens []string) string {
	formatted := ""

	NEWLINE_BEFORE := []string{"WHERE", "ORDER"}
	SPACE_NOT_ON := []string{"."}
	SPACE_NOT_BEFORE := []string{".", ";"}

	for idx, v := range tokens {

		/// next, prev = updatePrevNext()
		// update vars
		next := "" // clear in function ^
		// prev = "" // clear in function ^
		if idx < len(tokens)-1 {
			next = tokens[idx+1]
		}
		// if idx > 0 {
		// 	prev = tokens[idx-1]
		// }
		curr := tokens[idx]

		/// addSpace()
		willAddSpace := true
		if inside(next, SPACE_NOT_BEFORE) {
			willAddSpace = false
		}
		if inside(curr, SPACE_NOT_ON) {
			willAddSpace = false
		}
		// add space after all but last token
		if idx == len(tokens)-1 {
			willAddSpace = false
		}

		/// addNewline()
		willAddNewline := false
		// if next is WHERE or, newline
		if inside(next, NEWLINE_BEFORE) {
			willAddNewline = true
		}

		// print token (token itself is always printed!)
		formatted += v

		// formatting decisions
		if willAddSpace {
			formatted += " "
		}
		if willAddNewline {
			formatted += "\n"
		}
	}

	return formatted
}
