package main

import (
	"log"
	"os"
	"strings"

	"01.alem.school/git/ashagiro/ascii-art/funcs"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		log.Println("\nUsage: go run . [OPTION] [STRING]/[BANNER](depends on project) \n\nEX: go run . --<flagName>=<type> something standard")
		return
	}
	switch len(os.Args) {
	case 2:
		if strings.HasPrefix(os.Args[1], "--reverse") {
			funcs.Reverse()
		} else {
			funcs.Default()
		}
	case 3:
		if strings.HasPrefix(os.Args[1], "--color=") {
			funcs.Color()
		} else {
			if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {
				funcs.Default()
			} else {
				log.Println("\nUsage: go run . [OPTION] [STRING]/[BANNER] \n\n   EX: go run . something standard or shadow or thinkertoy \nOR EX: go run . --color=<color> something")
				return
			}
		}
	case 4:
		if strings.HasPrefix(os.Args[1], "--output=") {
			funcs.Output()
		} else if strings.HasPrefix(os.Args[1], "--color=") {
			funcs.Color()
		} else {
			log.Println("\nUsage: go run . [OPTION] [STRING]/[BANNER] \n\nOR EX: go run . --output=<fileName.txt> something standard \nOR EX: go run . --color=<color> <letters to be colored> something")
			return
		}
	}
}
