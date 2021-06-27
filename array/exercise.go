package main

import "fmt"

func main() {
	word := "golang is difficult"

	var elements []string
	var elCount []int

	for _, w := range word {
		letter := fmt.Sprintf("%c", w)
		if !isExist(letter, elements) {
			elements = append(elements, letter)
			elCount = append(elCount, count(letter, word))
		}
	}

	for i := 0; i < len(elements); i++ {
		fmt.Printf("letter: %s count: %d\n", elements[i], elCount[i])
	}
}

func count(letter string, word string) int {
	count := 0
	for _, w := range word {
		l := fmt.Sprintf("%c", w)
		if l == letter {
			count ++
		}
	}
	return count
}

func isExist(letter string, elements []string) bool {
	for _, e := range elements {
		if letter == e {
			return true
		}
	}
	return false
}