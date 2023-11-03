package main

import (
	"fmt"
)

// 定义一个Person结构体
type Person struct {
	Name string
	Age  int
}

// 定义一个接收者为Person类型的方法，用来获取Person的详细信息
func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// 定义一个接收者为指针类型的方法，用来修改Person的Age
func (p *Person) SetAge(age int) {
	p.Age = age
}

func main() {
	// 创建一个Person对象
	p := Person{Name: "Alice", Age: 25}

	// 调用Person的GetInfo方法，获取详细信息并打印
	fmt.Println(p.GetInfo())

	// 调用Person的SetAge方法，修改年龄为30
	p.SetAge(30)

	// 再次调用Person的GetInfo方法，查看年龄是否修改成功
	fmt.Println(p.GetInfo())
}
