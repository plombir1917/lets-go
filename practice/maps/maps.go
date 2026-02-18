package maps

import "fmt"

var m = make(map[string]int)

func Execute() {
	fmt.Println("\n___Maps___")
	wordsCounter("1")
	fmt.Println(m)

	wordsCounter("1")
	fmt.Println(m)

	wordsCounter("1")
	fmt.Println(m)
}

func wordsCounter(w string) {
	if v, ok := m[w]; ok {
		m[w] = v + 1
		return
	}
	m[w] = 1
}
