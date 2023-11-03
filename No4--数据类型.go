/*
这个示例中，包含了布尔型（boolean），整型（integer），无符号整型（unsigned），
浮点型（floatNum），字符串（str），字符（char），数组（arr），切片（slice），
映射（mapData），结构体（person），指针（pointer），函数（function），
接口（interfaceData），通道（channel）和空类型（emptyData）等各种数据类型。
运行此程序会输出相应数据类型的值。
*/
package main

import (
	"fmt"
)

func main() {
	// 布尔型
	var boolean bool = true
	fmt.Println("布尔型:", boolean)

	// 整型
	var integer int = 10
	fmt.Println("整型:", integer)

	// 无符号整型
	var unsigned uint = 20
	fmt.Println("无符号整型:", unsigned)

	// 浮点型
	var floatNum float32 = 3.14
	fmt.Println("浮点型:", floatNum)

	// 字符串
	var str string = "Hello, golang!"
	fmt.Println("字符串:", str)

	// 字符
	var char byte = 'A'
	fmt.Println("字符:", char)

	// 数组
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("数组:", arr)

	// 切片
	var slice []int = []int{1, 2, 3}
	fmt.Println("切片:", slice)

	// 映射
	var mapData map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("映射:", mapData)

	// 结构体
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"Alice", 25}
	fmt.Println("结构体:", person)

	// 指针
	var pointer *int
	pointer = &integer
	fmt.Println("指针:", *pointer)

	// 函数
	function := func(a int, b int) int {
		return a + b
	}
	result := function(2, 3)
	fmt.Println("函数:", result)

	// 接口
	var interfaceData interface{} = 3.14
	fmt.Println("接口:", interfaceData)

	// 通道
	channel := make(chan int)
	go func() {
		channel <- 1
	}()
	data := <-channel
	fmt.Println("通道:", data)

	// 空类型
	var emptyData interface{}
	fmt.Println("空类型:", emptyData)
}
