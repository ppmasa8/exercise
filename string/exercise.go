package main

import "fmt"

func main() {
	var word string
	var b, a rune
	fmt.Println("Please input word and letters ro convert. ex) papa p m\n> ")
	fmt.Scanf("%s %c %c\n", &word, &b, &a)

	for _, w := range word {
		if w == b {
			w = a
		}
		fmt.Printf("%c", w)
	}
}