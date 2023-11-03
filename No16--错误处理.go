/*
在上面的代码中，我们定义了一个名为divide的函数，该函数接受两个浮点数参数a和b，并返回a除以b的结果。如果b为零，则返回错误。
我们使用errors.New函数创建了一个新的错误。

在主函数中，我们调用divide函数并检查其结果和错误。如果err不为空，则打印错误消息，并使用return语句提前结束程序。
否则，打印结果。这种方式称为“提前返回”错误处理模式。
*/
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
