package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var StartTime = time.Now()
var NEWLINE_BEFORE = []string{"WHERE", "ORDER"}
var NEWLINE_ON = []string{";"}
var SPACE_NOT_ON = []string{"."}
var SPACE_NOT_BEFORE = []string{".", ";"}
var INCREMENT_INDENT_ON = []string{"SAVEPOINT"}
var DECREMENT_INDENT_BEFORE = []string{"ROLLBACK"}

var INDENTATION_LEVEL = 0
var INDENTATION_TOKEN = "    "

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tokens := tokenizeInput(os.Stdin)
	formatted := formatTokens(tokens)

	fmt.Print(formatted)
}

func formatTokens(tokens []string) string {
	formatted := ""

	INDENTATION_LEVEL = 0

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
		if inside(next, NEWLINE_BEFORE) {
			willAddNewline = true
		}
		if inside(curr, NEWLINE_ON) {
			willAddNewline = true
		}
		// if last token, add newline
		if idx == len(tokens)-1 {
			willAddNewline = true
		}

		// incIndent()
		if inside(curr, INCREMENT_INDENT_ON) {
			INDENTATION_LEVEL += 1
		}
		// decIndent() TODO... more lookahead
		if inside(next, DECREMENT_INDENT_BEFORE) {
			INDENTATION_LEVEL -= 1 // TODO error on < 0?
		}

		// print token (token itself is always printed!)
		formatted += v

		// formatting decisions
		if willAddNewline {
			formatted += "\n" + indent() // indent done together with newline
			willAddSpace = false         // newline always replaces space?
		}
		if willAddSpace {
			formatted += " "
		}
	}

	return formatted
}

func indent() string {
	return strings.Repeat(INDENTATION_TOKEN, INDENTATION_LEVEL)
}
