/*
这段代码定义了一个名为Person的结构体，它有三个字段：Name、Age和Address。
在main函数中，我们演示了如何创建结构体实例、访问和修改结构体字段的值，以及如何使用指针访问和修改结构体字段的值。
还演示了如何创建结构体指针和匿名结构体。
*/
package main

import (
	"fmt"
)

// 定义一个结构体
type Person3 struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// 创建一个Person结构体实例并初始化
	alice := Person3{Name: "Alice", Age: 20, Address: "No.1, Street"}

	// 访问结构体字段
	fmt.Println("Name:", alice.Name)
	fmt.Println("Age:", alice.Age)
	fmt.Println("Address:", alice.Address)

	// 修改结构体字段的值
	alice.Age = 21
	fmt.Println("Modified Age:", alice.Age)

	// 使用指针访问和修改结构体字段的值
	bob := &Person3{Name: "Bob", Age: 25, Address: "No.2, Street"}
	fmt.Println("Name:", bob.Name)
	fmt.Println("Age:", bob.Age)
	fmt.Println("Address:", bob.Address)
	bob.Age = 26
	fmt.Println("Modified Age:", bob.Age)

	// 创建结构体指针
	john := new(Person3)
	john.Name = "John"
	john.Age = 30
	john.Address = "No.3, Street"
	fmt.Println("Name:", john.Name)
	fmt.Println("Age:", john.Age)
	fmt.Println("Address:", john.Address)

	// 匿名结构体
	profile := struct {
		Username string
		Email    string
	}{
		Username: "user",
		Email:    "user@example.com",
	}
	fmt.Println("Username:", profile.Username)
	fmt.Println("Email:", profile.Email)
}
