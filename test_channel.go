package main

import "fmt"

func main() {

	intChan := make(chan int)

	fmt.Printf("intChan的值：%v, intChan本身的地址：%p\n", intChan, &intChan)
}
