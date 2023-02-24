package funcs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Default() {
	versions := []string{"shadow", "thinkertoy", "standard"}
	str := os.Args[1]
	for _, value := range str {
		if !(value > 0 && value <= 255) {
			fmt.Println("Only ascii chars!")
			return
		}
	}
	if len(os.Args) > 3 {
		ErrorMsg()
		return
	}
	version := "standard"
	if len(os.Args) == 3 {
		version = os.Args[2]
	}
	for _, value := range versions {
		if value == version {
			validFile := FileCheck(value)
			if !validFile {
				log.Println("Not a valid file!")
				os.Exit(1)
			}
			file, _ := os.Open(value + ".txt")
			defer file.Close()
			reader := bufio.NewScanner(file)
			art := ""
			for reader.Scan() {
				art += reader.Text() + "\n"
			}
			arr := MakeArt(str, []byte(art))
			fmt.Println(Write(arr))
			return
		}
	}
	fmt.Println("Invalid format! --> ", version)
}
