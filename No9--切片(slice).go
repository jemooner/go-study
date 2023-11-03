/*
注意：Go语言中的切片是指向底层数组的指针，因此对切片的修改也会影响到原始数组。
此外，切片的容量和长度可以使用内置的`len()`和`cap()`函数来获取。
*/
package main

import "fmt"

func main() {
	// 创建一个切片，包含 5 个元素，初始值都为 0
	nums := make([]int, 5)
	fmt.Println("初始切片：", nums)

	// 设置切片的值
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50
	fmt.Println("设置值后的切片：", nums)

	// 获取切片的长度和容量
	fmt.Println("切片的长度：", len(nums))
	fmt.Println("切片的容量：", cap(nums))

	// 截取切片
	slice := nums[1:4]
	fmt.Println("截取的切片：", slice)

	// 对切片进行追加
	slice = append(slice, 60)
	fmt.Println("追加后的切片：", slice)
	fmt.Println("原始切片也被修改：", nums)

	// 创建一个切片，直接初始化元素
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("水果切片：", fruits)
}
