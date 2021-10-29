package main

import "fmt"

func main() {
	a := 111
	b := 222
	swapValue(&a, &b)
	fmt.Println("a =", a, "  b =", b)
}
func swapValue(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
