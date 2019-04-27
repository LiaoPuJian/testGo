package main

import (
	"fmt"
)

func main() {
	a := [3]string{"a", "1", 2: "0"}
	fmt.Println(a)
	b := [...]int{50: 1}

	fmt.Println(
		b)

	p1 := new([10]int)
	fmt.Println(p1)

	c := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	fmt.Println(c)

	var x = [...]int{9, 7, 15, 66, 28, 17, 1}
	len := len(x)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if x[i] < x[j] {
				temp := x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
	fmt.Println(x)
}
