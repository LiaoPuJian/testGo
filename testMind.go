package main

import "fmt"

func main (){
	sliceAndArr();
}



//猜想，既然slice是数组的引用，那么定义一个slice引用某个数组，
// 修改了slice的值之后，数组的值也会修改吗？
func sliceAndArr(){
	//定义一个数组
	intArr := [5]int{1, 2, 3, 4, 5}

	intSlice := intArr[1:3]

	fmt.Println(intArr, intSlice)

	fmt.Printf("数组的类型:%T，切片的类型:%T\n", intArr, intSlice)

	//输出数组各个值的地址
	fmt.Printf("%p ", &intArr[0])
	fmt.Printf("%p ", &intArr[1])
	fmt.Printf("%p ", &intArr[2])
	fmt.Printf("%p ", &intArr[3])
	fmt.Printf("%p ", &intArr[4])
	fmt.Println()

	//输出切片各个值的地址
	fmt.Printf("%p ", &intSlice[0])
	fmt.Printf("%p ", &intSlice[1])
	fmt.Println()

	//这里可以看到，切片的第一个元素的地址和数组的第二个元素的地址是一样的
	//切片的第二个元素和数组的第三个元素的地址也是一样的，可以证实切片确实是数组的引用

	//这里修改切片的第一个值，看是否影响数组
	intSlice[0] = 10
	fmt.Println(intArr, intSlice)
	//可以看到确实影响了数组。

	//切片还可以继续切片，此时第二次产生的切片是否也指向第一次切片引用的数组？、
	intSlice2 := intSlice[:]
	//输出切片各个值的地址
	fmt.Printf("%p ", &intSlice2[0])
	fmt.Printf("%p ", &intSlice2[1])
	fmt.Println()
	//修改切片2的值，观察数组和切片1的值，可以看到数组和两个切片的值都变更了
	intSlice2[0] = 20
	fmt.Println(intArr, intSlice, intSlice2)

	//append函数的一些想法，声明一个切片，长度为3，容量为6
	intSlice3 := make([]int, 3, 6)
	fmt.Printf("%p ", &intSlice3[0])
	fmt.Printf("%p ", &intSlice3[1])
	fmt.Printf("%p ", &intSlice3[2])
	fmt.Println()

	//这个时候append一个小于等于三位的切片，观察地址变化，发现前三位的地址并未发生变化
	intSlice3 = append(intSlice3, intSlice2...)
	fmt.Printf("%p ", &intSlice3[0])
	fmt.Printf("%p ", &intSlice3[1])
	fmt.Printf("%p ", &intSlice3[2])
	fmt.Printf("%p ", &intSlice3[3])
	fmt.Printf("%p ", &intSlice3[4])
	fmt.Println()

	//此时再次追加一个切片，发现所有的地址都变化了，证明如果给切片追加超过其容量的值后，
	//系统会在内存中重新分配一块地址，并将旧地址的值copy到新的地址上去。
	intSlice3 = append(intSlice3, intSlice...)
	fmt.Printf("%p ", &intSlice3[0])
	fmt.Printf("%p ", &intSlice3[1])
	fmt.Printf("%p ", &intSlice3[2])
	fmt.Printf("%p ", &intSlice3[3])
	fmt.Printf("%p ", &intSlice3[4])
	fmt.Printf("%p ", &intSlice3[5])
	fmt.Printf("%p ", &intSlice3[6])
	fmt.Println()

	//如果不使用append而是直接增加值呢？   直接增加值的情况只要是超过当前数组的长度，就会提示越界
/*	intSlice4 := make([]int, 2, 3)
	fmt.Printf("%p ", &intSlice4[0])
	fmt.Printf("%p ", &intSlice4[1])
	intSlice4[2] = 20
	fmt.Printf("%p ", &intSlice4[0])
	fmt.Printf("%p ", &intSlice4[1])
	fmt.Printf("%p ", &intSlice4[2])
	fmt.Println()*/

	//copy函数猜想
	floatSlice := []float64{1.1, 1.2, 1.3}
	var floatSlice2 = make([]float64, 5)
	copy(floatSlice2, floatSlice)
	//使用copy函数后，查看对应的地址和值
	fmt.Println(floatSlice, floatSlice2)
	fmt.Printf("%p ", &floatSlice[0])
	fmt.Printf("%p ", &floatSlice[1])
	fmt.Printf("%p ", &floatSlice[2])
	fmt.Println()
	fmt.Printf("%p ", &floatSlice2[0])
	fmt.Printf("%p ", &floatSlice2[1])
	fmt.Printf("%p ", &floatSlice2[2])
	//发现地址完全不一样，证明使用了copy将一个切片的值给另一个切片后，两个切片之间是值copy关系
	//并未指向同一内存地址
}