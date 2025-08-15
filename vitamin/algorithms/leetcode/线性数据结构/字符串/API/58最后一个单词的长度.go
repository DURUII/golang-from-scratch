package main

import "strings"

func lengthOfLastWord(s string) int {
	// split with consecutive white space chars
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
