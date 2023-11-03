/*
这个示例代码中定义了一个worker函数，用于模拟一些工作。在主函数中，使用一个sync.WaitGroup来等待所有的goroutine完成。

在主函数中的for循环中，通过调用`go worker()`来创建并启动一个goroutine。每个goroutine会执行worker函数的内容。

最后，通过调用`wg.Wait()`等待所有的goroutine完成。一旦所有的goroutine都完成了它们的工作，程序会打印"All workers finished"的信息。

这是一个简单的示例，其中只包含了goroutine的基本使用方式。在实际的应用中，可能会涉及到更复杂的场景，比如使用channel来进行goroutine之间的通信，或者使用select语句来处理多个channel的情况等。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(1 * time.Second) // 模拟一些工作

	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有goroutine完成

	fmt.Println("All workers finished")
}
