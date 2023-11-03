/*
以下代码包含了常见的循环语句示例：

1. 示例1演示了使用`for`关键字循环执行一段代码块的基本用法，打印出0到4的数字。
2. 示例2演示了使用`range`关键字遍历一个切片，并打印出索引和对应值。
3. 示例3演示了使用无限循环和`break`语句来实现一个简单的计数器，打印出1到5的数字。
4. 示例4演示了循环嵌套的用法，通过两层循环打印出1到3的所有组合。
*/
package main

import "fmt"

func main() {
	// 示例1：for循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 示例2：for range循环
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("索引：%d，值：%d\n", index, value)
	}

	// 示例3：无限循环
	count := 0
	for {
		count++
		if count > 5 {
			break
		}
		fmt.Println(count)
	}

	// 示例4：循环嵌套
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
