package main

import (
	"os"
	"text/scanner"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func inside(e string, s []string) bool {
	return contains(s, e)
}

func tokenizeInput() []string {
	var tokens []string

	var s scanner.Scanner
	s.Init(os.Stdin)
	var token rune
	token = s.Scan()
	for token != scanner.EOF {
		tokens = append(tokens, s.TokenText())
		token = s.Scan()
	}

	return tokens
}
