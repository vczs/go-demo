package main

import "fmt"

func main() {
	a := 111
	b := 222
	a, b = swapValue(a, b)
	fmt.Println("a =", a, "  b =", b)
}
func swapValue(a int, b int) (int, int) {
	a = a + b // a=a+b
	b = a - b // b=a+b-b=a即b=a
	a = a - b // a=a+b-a=b即a=b
	//此时a=b,b=a即a和b交换了值
	return a, b
}
