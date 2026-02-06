package defer_example

import "fmt"

var global = 5

func DeferForGlobal() {
	defer func(local int) {
		global = local
		fmt.Println(global)
	}(global)
	global = 10
	fmt.Println(global)
}
