/*
以下代码定义了两个自定义类型：`Person`和`MyInt`。`Person`类型有两个字段：`Name`和`Age`，以及一个方法`SayHello()`。

`MyInt`类型继承自内置类型`int`。在主函数中，我们创建了一个`Person`类型的变量并调用了其方法，创建了一个`MyInt`类型的变量并传递给了函数进行计算。

请注意，`type`关键字可以用来创建自定义类型、方法接收器或函数参数/返回值的类型别名，这些是Go语言中type使用的常见用法。
*/
package main

import "fmt"

// 定义一个自定义类型Person
type Person2 struct {
	Name string
	Age  int
}

// 定义一个自定义类型MyInt
type MyInt int

// 定义一个方法，接收Person类型的参数
func (p Person2) SayHello() {
	fmt.Printf("Hello, my name is %s!\n", p.Name)
}

// 定义一个函数，接收MyInt类型的参数
func AddOne(num MyInt) MyInt {
	return num + 1
}

func main() {
	// 创建一个Person类型的变量
	person := Person2{
		Name: "Alice",
		Age:  25,
	}

	// 调用Person类型的方法
	person.SayHello()

	// 创建一个MyInt类型的变量
	num := MyInt(5)

	// 调用函数，传递MyInt类型的参数
	result := AddOne(num)

	// 输出结果
	fmt.Printf("The result is: %d\n", result)
}
