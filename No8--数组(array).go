/*
下面是一个示例代码，演示了如何在Go语言中创建数组、访问数组元素和遍历数组
*/
package main

import "fmt"

func main() {
	// 创建一个长度为 5 的整数数组
	var arr [5]int

	// 给数组赋值
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// 访问数组元素
	fmt.Println(arr[0]) // 输出: 10
	fmt.Println(arr[3]) // 输出: 40

	// 遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 使用字面量创建数组
	arr2 := [3]string{"foo", "bar", "baz"}
	fmt.Println(arr2[0]) // 输出: foo
	fmt.Println(arr2[2]) // 输出: baz
}
