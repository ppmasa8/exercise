package main

import "fmt"

func main() {
	a, b, c := 1, 1, 1

	a *= 2
	fmt.Printf("main() aの値: %d\n\n", a)

	multiply2(b)
	fmt.Printf("main() bの値: %d, bのアドレス: %p\n\n", b, &b)

	multiply2pointer(&c)
	fmt.Printf("main() cの値: %d, cのアドレス: %p\n\n", c, &c)
}

func multiply2(b int) {
	b *= 2
	fmt.Printf("multiply2() bの値: %d, bのアドレス: %p\n", b, &b)
}

func multiply2pointer(c *int) {
	*c *= 2
	fmt.Printf("multiply2pointer() cの値: %d, cのアドレス: %p\n", *c, c)
}