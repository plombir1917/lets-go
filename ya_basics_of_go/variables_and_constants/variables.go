package main

import "fmt"

func main() {
    // определите переменные ver, id, pi
    
    // ver, id, pi := "v0.0.1", 0, 3.1415 короткая нотация в одну строку
    // var ver, id, pi = "v0.0.1", 0, 3.1415 объявление без указания типа в одну строку

    // полная запись с указанием типа и начального значения
    // var ver string = "v0.0.1"
    // var id int = 0
    // var pi float32 = 3.1415

    // объявление схожих по смыслу переменных в блоке
    var(
        height, length int
    )
    fmt.Println(height, length)

    // объявление без указания типа, пример из решебника
    var ver = "v0.0.1"
    var id int
    var pi = 3.1415


    fmt.Println("ver =", ver, "id =", id, "pi =", pi)
}
