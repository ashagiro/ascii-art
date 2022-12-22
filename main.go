package main

import (
	"os"

	"01.alem.school/git/ashagiro/ascii-art/functions"
)

func main() {
	if !(len(os.Args) >= 2 && len(os.Args) <= 4) {
		functions.ErrorMsg()
		return
	}
	// Here program checks if there are COMMAND OPTIONs present
	if functions.Option(os.Args[1]) {
	} else {
		// if no commands found default ascii-art project ([STRING] [BANNER] or [STRING]) will operate
		functions.Default()
	}
}
