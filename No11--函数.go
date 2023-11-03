package main

import "fmt"

// 定义一个函数，接受两个整数参数，并返回它们的和
func add(x int, y int) int {
	return x + y
}

// 定义一个函数，接受一个函数作为参数，并调用该函数
func apply(f func(int, int) int, x int, y int) int {
	return f(x, y)
}

// 定义一个函数，返回一个函数
func getMultiplier() func(int) int {
	return func(x int) int {
		return x * 2
	}
}

func main() {
	// 调用 add 函数并打印结果
	sum := add(1, 2)
	fmt.Println("1 + 2 =", sum)

	// 使用 apply 函数调用 add 函数并打印结果
	appliedSum := apply(add, 3, 4)
	fmt.Println("3 + 4 =", appliedSum)

	// 调用 getMultiplier 函数获取一个乘以 2 的函数，并使用该函数计算结果并打印
	multiplier := getMultiplier()
	result := multiplier(5)
	fmt.Println("5 * 2 =", result)
}
