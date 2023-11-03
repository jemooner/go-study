/*
这个示例中，首先声明了一个整型变量`num`并赋值为42。然后，使用`&`运算符获取`num`的内存地址，并将其赋值给了一个整型指针变量`ptr`。

然后，通过`fmt.Println()`函数进行输出，展示`num`的值、内存地址以及`ptr`的值。

接着使用`*`运算符访问指针所指向的变量的值，输出了`ptr`指向的变量的值。

最后，通过修改`*ptr`的值，同时也会修改`num`的值。

这个示例展示了如何声明指针变量、获取指针所指向变量的值以及修改指针所指向变量的值。
*/
package main

import "fmt"

func main() {
	// 声明一个整型变量
	var num int = 42

	// 声明一个指向整型的指针变量，并将其指向num的内存地址
	var ptr *int = &num

	fmt.Println("Value of num:", num)    // 输出: Value of num: 42
	fmt.Println("Address of num:", &num) // 输出: Address of num: 0xc000014090
	fmt.Println("Value of ptr:", ptr)    // 输出: Value of ptr: 0xc000014090

	// 通过指针访问变量的值，使用*运算符
	fmt.Println("Value stored in ptr:", *ptr) // 输出: Value stored in ptr: 42

	// 修改指针所指向的变量的值
	*ptr = 100
	fmt.Println("Modified value of num:", num) // 输出: Modified value of num: 100
}
