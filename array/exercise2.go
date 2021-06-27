package main

import "fmt"

func main() {
	word := "hello world"

	elements := make(map[string]int)
	for _, w := range word {
		key := fmt.Sprintf("%c", w)
		if _, ok := elements[key]; ok {
			elements[key]++
		} else {
			elements[key] = 1
		}
	}

	for k, v := range elements {
		fmt.Printf("letter: %s count: %d\n", k, v)
	}
	fmt.Println(elements)
}