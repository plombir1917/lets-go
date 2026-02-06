package main

import "fmt"

func secondExercise() {
	pricelist := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	total := 0
	fmt.Println("Перечень деликатесов:")
	for k, v := range pricelist {
		if v > 500 {
			fmt.Println(k)
		}
	}
	for _, v := range order {
		total += pricelist[v]
	}
	fmt.Println("Стоимость заказа ", total)

}

func main() {
	m := make(map[string]string)
	m["foo"] = "bar"
	m["bazz"] = "yup"
	for k, v := range m {
		// k будет перебирать ключи
		// v – соответствующие ключам значения
		fmt.Printf("Ключ %v, имеет значение %v \n", k, v)
	}

	secondExercise()
}
