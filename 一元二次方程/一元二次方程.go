package main

import(
	"fmt"
	"math"
)

func main(){
	var a = 3.0
	var b float64 = 100.0
	c := 6.0
	m := b * b - 4 * a * c
	if m > 0 {
		x1 := (-b + math.Sqrt(m))/2*a
		x2 := (-b - math.Sqrt(m))/2*a
		fmt.Printf("有两个解：\nx1 = %v\nx2 = %v\n",x1,x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m))/2*a
		fmt.Printf("有一个解：\nx1 = %v\n",x1)
	} else {
		fmt.Println("无解...")
	}
}
