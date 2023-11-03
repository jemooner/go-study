/*
首先我们导入了需要的包在这个示例代码中，首先我们导入了需要的包。然后，我们创建一个`bufio.Reader`对象来从标准输入流中读取输入。
然后，我们使用`Printf`函数向用户提示输入他们的名字，并使用`ReadString`方法从标准输入流中读取用户输入的一行文本。
最后，我们使用`Printf`函数再次将用户输入的名字打印出来。

你可以将此示例代码保存到一个名为`main.go`的文件中，然后使用`go run main.go`命令来运行它。
在控制台中，你将看到一个提示要求你输入你的名字，然后你可以输入你的名字并按下回车键。程序将打印出你输入的名字。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开标准输入流
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入你的名字：")

	// 从标准输入流中读取一行文本
	name, _ := reader.ReadString('\n')

	fmt.Printf("你的名字是：%s", name)
}
