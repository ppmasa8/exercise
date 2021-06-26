package main

import "fmt"

func main() {
	a, b, c := 1, 3, 5
	swap1(a ,b ,c)
	fmt.Println(a, b, c)

	swap2(&a ,&b ,&c)
	fmt.Println(a, b, c)
}

func swap1(a, b, c int) {
	a, b, c = b, c, a
}

func swap2(a, b, c *int) {
	*a, *b, *c = *b, *c, *a
}