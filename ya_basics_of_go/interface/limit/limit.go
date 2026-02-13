package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func LimitReader(r io.Reader, n int) io.Reader {
	bytes, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	limitedBytes := bytes[:n]

	return strings.NewReader(string(limitedBytes))
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read/n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)

	if err != nil {
		log.Fatal(err)
	}
}
