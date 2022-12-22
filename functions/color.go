package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type colors struct {
	color string
	code  string
}

// Following link was helpful: https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html

func Color(s string) {
	array := []colors{
		{"black", "\u001b[30;1m"}, {"yellow", "\u001b[33;1m"}, {"red", "\u001b[31;1m"}, {"green", "\u001b[32;1m"}, {"blue", "\u001b[34;1m"}, {"magenta", "\u001b[35;1m"}, {"cyan", "\u001b[36;1m"}, {"white", "\u001b[37;1m"}, {"purple", "\033[95m"}, {"orange", "\033[38;2;255;165;0m"},
	}
	paint := ""
	for _, v := range array {
		if v.color == s {
			paint = v.code
		}
	}
	if len(paint) == 0 {
		ErrorMsgColor()
		return
	}
	validFile := FileCheck("standard")
	if !validFile {
		fmt.Println("Not a valid file!")
		return
	}
	file, _ := os.Open("standard.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	art := ""
	for reader.Scan() {
		art += reader.Text() + "\n"
	}
	if len(os.Args) == 3 {
		str := os.Args[2]
		arr := MakeArt(str, []byte(art))
		fmt.Println(paint + Write(arr) + "\u001b[0m")
		return
	}
	if len(os.Args) == 4 {
		str := os.Args[3]
		arr := MakeArt(str, []byte(art))
		fmt.Println(Write2(arr, paint))
		return
	}
}

func Write2(arr [][]string, paint string) string {
	var word string
	startP := 0
	str := RewriteWord(os.Args[3])
	if !LettersPresent(str) {
		return "ERROR: letters are not present in the word"
	}
	for index, value := range arr {
		if len(value) == 1 && startP == index {
			word += ("\n")
			startP += 1
		} else if len(value) == 1 {
			for i := 1; i <= 8; i++ {
				for j := startP; j < index; j++ {
					if string(str[j]) != "\n" && strings.Contains(os.Args[2], string(str[j])) {
						word += paint + string(arr[j][i]) + "\u001b[0m"
					} else {
						word += string(arr[j][i])
					}
				}
				word += "\n"
			}
			startP = index + 1
		} else if index+1 == len(arr) {
			for i := 1; i <= 8; i++ {
				for j := startP; j < index+1; j++ {
					if os.Args[2] != "\n" && strings.Contains(os.Args[2], string(str[j])) {
						word += paint + string(arr[j][i]) + "\u001b[0m"
					} else {
						word += string(arr[j][i])
					}
				}
				word += "\n"
			}
		}
	}
	return word[:len(word)-1]
}

func RewriteWord(arg string) string {
	word := ""
	for i := 0; i < len(arg); i++ {
		if i != len(arg)-1 && string(arg[i]) == "\\" && string(arg[i+1]) == "n" {
			word += "\n"
			i = i + 1
		} else {
			word += string(arg[i])
		}
	}
	return word
}

func LettersPresent(word string) bool {
	check := true
	letters := os.Args[2]
	for _, k := range letters {
		if !strings.Contains(word, string(k)) {
			check = false
		}
	}
	return check
}
