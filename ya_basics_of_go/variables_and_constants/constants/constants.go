package main

import "fmt"

const (
    one = 1 + iota * 2
    three
    five
    seven
    nine
    eleven
)

func main() {
    fmt.Println(one, three, five, seven, nine, eleven)
}