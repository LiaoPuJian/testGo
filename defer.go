package main

import (
	"fmt"
)

func main(){
	fmt.Println("Start")

	defer fmt.Println("a1")

	defer fmt.Println("a2")

	defer fmt.Println("a3")


	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("End")
}
