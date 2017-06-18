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
var NEWLINE_ON = []string{";", "BEGIN", "EXCEPTION", "THEN"}
var SPACE_NOT_ON = []string{".", "("}
var SPACE_NOT_BEFORE = []string{".", ";", ",", "(", ")"}
var INCREMENT_INDENT_ON = []string{"SAVEPOINT", "BEGIN", "THEN"}
var DECREMENT_INDENT_BEFORE = []string{"ROLLBACK", "END", "EXCEPTION"}

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

		/// next, next2 = updateLookaheads()
		// update vars
		prev := ""
		prev2 := ""
		prev3 := ""
		next := ""  // clear in function ^
		next2 := "" // clear in function ^
		if idx > 0 {
			prev = tokens[idx-1]
		}
		if idx > 1 {
			prev2 = tokens[idx-2]
		}
		if idx > 2 {
			prev3 = tokens[idx-3]
		}
		if idx < len(tokens)-1 {
			next = tokens[idx+1]
		}
		if idx < len(tokens)-2 {
			next2 = tokens[idx+2]
		}
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
		// dollar quoting
		if curr == "$" && next2 == "$" {
			willAddSpace = false
		}
		if prev == "$" && next == "$" {
			willAddSpace = false
		}

		/// addNewline()
		willAddNewline := false
		if inside(next, NEWLINE_BEFORE) {
			willAddNewline = true
		}
		if prev != "RAISE" && inside(curr, NEWLINE_ON) {
			willAddNewline = true
		}
		// if last token, add newline
		if idx == len(tokens)-1 {
			willAddNewline = true
		}
		// dollar quoting
		if prev3 == "DO" && prev2 == "$" && curr == "$" {
			willAddNewline = true
		}

		// incIndent()
		if prev2 != "ROLLBACK" && inside(curr, INCREMENT_INDENT_ON) {
			INDENTATION_LEVEL += 1
		}
		// decIndent() TODO... more lookahead
		if curr != "RAISE" && inside(next, DECREMENT_INDENT_BEFORE) {
			INDENTATION_LEVEL -= 1 // TODO error on < 0?
			if INDENTATION_LEVEL < 0 {
				log.Println("INDENTATION_LEVEL was negative!")
				INDENTATION_LEVEL = 0
			}
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
