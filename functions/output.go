package functions

import (
	"bufio"
	"fmt"
	"os"
)

func Output(s string) {
	versions := []string{"shadow", "thinkertoy", "standard"}
	str := os.Args[2]
	for _, value := range str {
		if !(value > 0 && value <= 255) {
			fmt.Println("Only ascii chars!")
			return
		}
	}
	version := "standard"
	if len(os.Args) == 4 {
		version = os.Args[3]
	}
	for _, value := range versions {
		if value == version {
			validFile := FileCheck(value)
			if !validFile {
				fmt.Println("Not a valid file!")
				return
			}
			file, _ := os.Open(value + ".txt")
			defer file.Close()
			reader := bufio.NewScanner(file)
			art := ""
			for reader.Scan() {
				art += reader.Text() + "\n"
			}
			arr := MakeArt(str, []byte(art))

			file, err := os.Create(s)
			if err != nil {
				fmt.Println("Oopps! File " + s + " can not be created.")
				return
			}
			defer file.Close()

			u, err := file.WriteString(Write(arr) + "\n\n")
			Error(err)
			u = 1 + u
			return
		}
	}
	fmt.Println("Invalid format! --> ", version)
}

func Error(err error) {
	if err != nil {
		fmt.Println("Oopps! Can not save the text.")
		return
	}
}
