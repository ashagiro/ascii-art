package funcs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// useful links:
// https://devtidbits.com/2021/09/15/center-text-in-a-terminal-with-go/
// https://www.dotnetperls.com/padding-go

func Align(s string) {
	str := os.Args[2]
	version := os.Args[3]
	if !(s == "center" || s == "left" || s == "right" || s == "justify") {
		log.Println("Incorrect format >> " + s)
		os.Exit(1)
	}
	if !(version == "thinkertoy" || version == "shadow" || version == "standard") {
		log.Println("Incorrect format >> " + version)
		os.Exit(1)
	}
	// following lines are to calculate the width of the terminal window
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	arr := []string{}
	arr = strings.Split(string(out[:len(out)-1]), " ")
	width, err := strconv.Atoi(arr[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	validFile := FileCheck(version)
	if !validFile {
		log.Println("Not a valid file!")
		os.Exit(1)
	}
	file, _ := os.Open(version + ".txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	art := ""
	for reader.Scan() {
		art += reader.Text() + "\n"
	}

	str = strings.ReplaceAll(str, "\\n", "\n")
	layers := strings.Split(str, "\n")
	for _, keys := range layers {
		array := strings.Split(keys, " ")
		artlen := []int{}
		for _, k := range array {
			artlen = append(artlen, len(Write(MakeArt(k, []byte(art))))/8)
		}
		word := Write(MakeArt(keys, []byte(art)))
		if len(word)/8 > width {
			log.Println("Too many words")
			os.Exit(1)
		}
		fmt.Println(ApplyFunc(artlen, s, word, width))
	}
}

func ApplyFunc(artlen []int, loc string, s string, width int) string {
	var str string
	arr := strings.Split(s, "\n")
	if len(arr) == 0 {
		return str
	}
	for i := 0; i < len(arr); i++ {
		if loc == "center" {
			str += Center(arr[i], (width - len(arr[i]) - 2))
		}
		if loc == "right" {
			str += Right(arr[i], (width - len(arr[i]) - 2))
		}
		if loc == "left" {
			str += Left(arr[i], (width - len(arr[i]) - 2))
		}
		if loc == "justify" {
			str += Justify(artlen, arr[i], (width - len(arr[i]) - 2))
		}
	}
	return str
}

func Right(s string, n int) string {
	return "|" + strings.Repeat(" ", n) + s + "|"
}

func Left(s string, n int) string {
	return "|" + s + strings.Repeat(" ", n) + "|"
}

func Center(s string, n int) string {
	div := n / 2
	return "|" + strings.Repeat(" ", div) + s + strings.Repeat(" ", n-div) + "|"
}

func Justify(artlen []int, s string, n int) string {
	arr := SplitWord(artlen, s)
	if len(arr) <= 1 {
		return "|" + s + strings.Repeat(" ", n) + "|"
	}
	div := n / (len(arr) - 1)
	str := string(arr[0])
	res := n - div*(len(arr)-1)
	for i := 1; i < len(arr); i++ {
		if i == len(arr)-1 {
			str += "      " + strings.Repeat(" ", div+res) + string(arr[i])
			break
		}
		str += "      " + strings.Repeat(" ", div) + string(arr[i])
	}
	return "|" + str + "|"
}

func SplitWord(n []int, s string) []string {
	arr := []string{}
	if len(n) < 2 {
		arr = append(arr, s)
		return arr
	}
	i := 0
	for _, v := range n {
		arr = append(arr, s[i:i+v])
		i += 6
		i = i + v
	}
	return arr
}
