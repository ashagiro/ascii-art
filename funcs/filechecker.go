package funcs

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
)

func FileCheck(filename string) bool {
	fileHashes := make(map[string]string)
	fileHashes["standard"] = "ac85e83127e49ec42487f272d9b9db8b"
	fileHashes["shadow"] = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	fileHashes["thinkertoy"] = "86d9947457f6a41a18cb98427e314ff8"
	data, err := os.ReadFile(filename + ".txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if !(fileHashes[filename] == GetMD5Hash(string(data))) {
		log.Println("File has been modified!")
		os.Exit(1)
	}
	return true
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ErrorMsgOutput() {
	log.Println("Usage: go run . [OPTION] [STRING] [BANNER]" + "\n\n" + "EX: go run . --output=<fileName.txt> something standard")
	os.Exit(1)
}

func ErrorMsgColor() {
	log.Println("Usage: go run . [OPTION] [STRING]" + "\n\n" + "EX: go run . --color=<color> <letters to be colored> something")
	os.Exit(1)
}

func ErrorMsg() {
	log.Println("Usage: go run . [STRING] [BANNER]" + "\n\n" + "EX: go run . something standard")
	os.Exit(1)
}

func ErrorMsgAlign() {
	log.Println("Usage: go run . [OPTION] [STRING] [BANNER]" + "\n\n" + "Example: go run . --align=right something standard")
	os.Exit(1)
}
