package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	NEWLINE_BEFORE := []string{"WHERE", "ORDER"}
	SPACE_NOT_ON := []string{"."}
	SPACE_NOT_BEFORE := []string{".", ";"}

	next := ""
	// prev := ""
	curr := ""
	tokens := tokenizeInput()
	for idx, v := range tokens {

		/// next, prev = updatePrevNext()
		// update vars
		next = "" // clear in function ^
		// prev = "" // clear in function ^
		if idx < len(tokens)-1 {
			next = tokens[idx+1]
		}
		// if idx > 0 {
		// 	prev = tokens[idx-1]
		// }
		curr = tokens[idx]

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

		// print it (character always printed)
		fmt.Print(v)

		// formatting decisions
		if willAddSpace {
			fmt.Print(" ")
		}
		if willAddNewline {
			fmt.Print("\n")
		}
	}
}
