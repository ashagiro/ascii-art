package funcs

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Reverse() {
	testfilename, err := ReverseFlag()
	if err != nil {
		fmt.Printf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>\n")
		os.Exit(0)
	}
	bannername := "./banners/standard.txt"
	// Check for hash
	if !FileCheck(bannername) {
		fmt.Println("Error: Wrong hash!!!\nThe file has been modified")
		return
	}
	// Get map of all characters from bannerfile
	MapOfBanner := GetMapOfBanner(bannername)
	// ArrStrOfBanner, err := funcs.GetArrStr(bannername)
	if err != nil {
		if err != nil {
			fmt.Println("Error: can't read file")
			os.Exit(0)
		}
	}

	ArrStrOfTestfile, err := GetArrStr(testfilename)
	if err != nil {
		fmt.Println("Error: can't read file")
		os.Exit(0)
	}

	CheckTestFile(ArrStrOfTestfile)

	for row := 0; row <= len(ArrStrOfTestfile)-1; row++ {
		if ArrStrOfTestfile[row] == "" {
			fmt.Println()
		} else {
			index := 0
			for i := 0; i <= len(ArrStrOfTestfile[row])-4; i++ {
				charstr := ""
				for j := row; j <= row+7; j++ {
					charstr += ArrStrOfTestfile[j][index:(i + 4)]
				}
				for key, val := range MapOfBanner {
					if charstr == val {
						fmt.Print(string(key))
						index += (i + 4) - index
						break
					}
				}
			}
			fmt.Println()
			row += 7
		}
	}
}

func GetMapOfBanner(filename string) map[rune]string {
	ArrStr, err := GetArrStr(filename)
	if err != nil {
		fmt.Println("Error: can't read file")
		os.Exit(0)
	}
	mymap := make(map[rune]string)
	var ArtOfChar string
	k := 32
	for i := 0; i < len(ArrStr); i++ {
		if i == len(ArrStr)-1 && ArtOfChar != "" {
			ArtOfChar += ArrStr[i]
			mymap[rune(k)] = ArtOfChar
		}
		if ArrStr[i] != "" {
			ArtOfChar += ArrStr[i]
		} else if len(ArtOfChar) != 0 {
			mymap[rune(k)] = ArtOfChar
			ArtOfChar = ""
			k++
		}
	}
	return mymap
}

func ReverseFlag() (string, error) {
	reverse := flag.String("reverse", "", "")
	flag.Usage = func() {
		fmt.Printf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>\n")
		os.Exit(0)
	}
	flag.CommandLine.Parse(os.Args[1:])
	check := regexp.MustCompile(`\w+.txt`)
	match := check.Match([]byte(*reverse))
	if !match {
		return "", errors.New("\nUsage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
	}
	return *reverse, nil
}

func CheckTestFile(ArrStrOfTestfile []string) {
	if ArrStrOfTestfile == nil {
		fmt.Println("ERROR: the testfile is empty")
		os.Exit(0)
	} else if len(ArrStrOfTestfile) < 8 {
		fmt.Println("ERROR: wrong testfile")
		os.Exit(0)
	}
	r := 0
	for i := 0; i <= len(ArrStrOfTestfile)-1; i++ {
		if r == 7 {
			i += r
			r = 0
		}
		if i < len(ArrStrOfTestfile) {
			if len(ArrStrOfTestfile[i]) >= 4 {
				for j := i; j <= i+7; j++ {
					if j != i+7 {
						if len(ArrStrOfTestfile[j]) != len(ArrStrOfTestfile[j+1]) {
							fmt.Println("ERROR: wrong testfile")
							os.Exit(0)
						}
					}
				}
				r = 7
			} else {
				if ArrStrOfTestfile[i] != "" {
					fmt.Println("ERROR: wrong testfile")
					os.Exit(0)
				}
			}
		} else {
			break
		}

	}
}

func GetArrStr(filename string) ([]string, error) {
	if strings.HasPrefix(os.Args[1], "--reverse=") {
	}
	var text []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		text = append(text, str)
	}
	err = file.Close()
	if err != nil || scanner.Err() != nil {
		return nil, err
	}
	return text, nil
}
