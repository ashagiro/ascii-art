package functions

func DefineAsciiSymbol() map[int]int {
	symbols := make(map[int]int)
	count := 2
	for i := 32; i < 127; i++ {
		symbols[i] = count
		count += 9
	}
	return symbols
}

func MakeArt(value string, data []byte) [][]string {
	symbolStartingPoint := DefineAsciiSymbol()
	asciiWordArr := [][]string{}
	for i := 0; i < len(value); i++ {
		subArr := []string{}
		count := 1
		word := ""

		if i != len(value)-1 && string(value[i]) == "\\" && string(value[i+1]) == "n" {
			subArr = append(subArr, "\n")
			i = i + 1
		} else if string(value[i]) == "\n" {
			subArr = append(subArr, "\n")
		} else {
			for _, a_char := range data {
				if a_char == '\n' {
					count++
				}
				if count >= symbolStartingPoint[int(value[i])] && count < symbolStartingPoint[int(value[i])]+9 {
					if a_char == '\n' {
						subArr = append(subArr, word)
						word = ""
					} else {
						word += string(a_char)
					}
				}
			}
		}
		asciiWordArr = append(asciiWordArr, subArr)

	}
	return asciiWordArr
}

func Write(arr [][]string) string {
	var str string
	startP := 0
	for index, value := range arr {
		if len(value) == 1 && startP == index {
			str += ("\n")
			startP += 1
		} else if len(value) == 1 {
			for i := 1; i <= 8; i++ {
				for j := startP; j < index; j++ {
					str += string(arr[j][i])
				}
				str += "\n"
			}
			startP = index + 1
		} else if index+1 == len(arr) {
			for i := 1; i <= 8; i++ {
				for j := startP; j < index+1; j++ {
					str += string(arr[j][i])
				}
				str += "\n"
			}
		}
	}
	if len(str) == 0 {
		return str
	}
	return str[:len(str)-1]
}
