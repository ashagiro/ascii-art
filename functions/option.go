package functions

import (
	"fmt"
	"os"
)

func Option(s string) bool {
	if len(s) < 10 {
		return false
	}
	if s[0:9] == "--output=" {
		if len(os.Args[1]) < 14 || len(os.Args) == 2 {
			ErrorMsgOutput()
			return true
		}
		_, err := os.ReadFile(s[9:])
		if err == nil {
			fmt.Println("Oopps! " + s[9:] + " file already exists.")
			return true
		}
		if os.Args[1][len(s)-4:len(s)] != ".txt" {
			fmt.Println("Oopps! Not a txt file.")
			return true
		}
		Output(os.Args[1][9:])
		return true
	}
	if s[0:8] == "--color=" {
		if !(len(os.Args) == 3 || len(os.Args) == 4) || len(os.Args[1]) < 10 {
			ErrorMsgColor()
			return true
		}
		Color(os.Args[1][8:])
		return true
	}
	if s[0:8] == "--align=" {
		if len(os.Args) != 4 {
			ErrorMsgAlign()
			return true
		}
		Align(os.Args[1][8:])
		return true
	}
	if s[0:10] == "--reverse=" {
		fmt.Println("Work in progress")
		return true
	}
	return false
}

func ErrorMsgOutput() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]" + "\n\n" + "EX: go run . --output=<fileName.txt> something standard")
}

func ErrorMsgColor() {
	fmt.Println("Usage: go run . [OPTION] [STRING]" + "\n\n" + "EX: go run . --color=<color> <letters to be colored> something")
}

func ErrorMsg() {
	fmt.Println("Usage: go run . [STRING] [BANNER]" + "\n\n" + "EX: go run . something standard")
}

func ErrorMsgAlign() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]" + "\n\n" + "Example: go run . --align=right something standard")
}
