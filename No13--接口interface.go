/*
代码中定义了一个Shape接口，包含了两个方法Area()和Perimeter()，然后定义了Rectangle和Circle两个结构体，
分别实现了Shape接口的方法。在main函数中，通过赋值给接口类型的变量，展示了接口的使用。首先将Rectangle对象赋值给shape，
调用Area()和Perimeter()方法计算矩形的面积和周长；然后将Circle对象赋值给shape，调用Area()和Perimeter()方法计算圆形的面积和周长。
通过这种方式，我们可以实现对不同结构体对象的统一操作。运行该代码，将输出矩形和圆形的面积和周长。
*/
package main

import (
	"fmt"
)

// 定义一个接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义一个圆形结构体
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {
	// 创建一个矩形对象
	rectangle := Rectangle{
		Width:  5,
		Height: 3,
	}

	// 创建一个圆形对象
	circle := Circle{
		Radius: 4,
	}

	// 接口类型的变量可以接收任何实现了该接口的对象
	var shape Shape

	// 将矩形对象赋值给接口类型的变量
	shape = rectangle
	fmt.Printf("矩形的面积为 %.2f，周长为 %.2f\n", shape.Area(), shape.Perimeter())

	// 将圆形对象赋值给接口类型的变量
	shape = circle
	fmt.Printf("圆形的面积为 %.2f，周长为 %.2f\n", shape.Area(), shape.Perimeter())
}
