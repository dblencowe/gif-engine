package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func Dump(data any) {
	b, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(b))
}

func DumpToFile(path string, data any) {
	b, _ := json.MarshalIndent(data, "", " ")
	os.WriteFile(path, b, os.ModePerm)
}