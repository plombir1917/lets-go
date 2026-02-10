package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	PrintAllFiles(".", "main")
}

func PrintAllFiles(path string, filter string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		filename := filepath.Join(path, f.Name())

		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}

		if f.IsDir() {
			PrintAllFiles(filename, filter)
		}
	}
}
