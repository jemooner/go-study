package main

import "fmt"

func main() {
	// 创建一个空的map
	m := make(map[string]int)

	// 添加键值对
	m["apple"] = 1
	m["banana"] = 2
	m["orange"] = 3

	// 打印map中的键值对
	fmt.Println("map:", m)

	// 获取map中某个键的值
	value := m["apple"]
	fmt.Println("apple:", value)

	// 删除map中的某个键值对
	delete(m, "banana")
	fmt.Println("map:", m)

	// 检查某个键是否存在map中
	_, isIn := m["orange"]
	fmt.Println("orange exists:", isIn)

	// 遍历map中的所有键值对
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
