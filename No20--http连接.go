/*
在这个示例中，我们首先创建了一个自定义的http.Transport，然后设置了最大连接数和最大空闲连接数。接下来，使用这个自定义的Transport创建了一个http.Client。

在main函数中，我们定义了要发送的URL列表，并且使用一个sync.WaitGroup来等待所有请求的完成。然后，我们通过循环遍历URL列表，为每个URL启动一个goroutine来发送请求。最后，我们调用Wait方法等待所有请求完成。

通过使用http长连接池，我们可以重复使用现有连接，避免了每次请求都创建和关闭连接的开销，提高了性能。
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var client *http.Client

func init() {
	// 创建一个自定义的http.Transport，设置最大连接数和最大空闲连接数
	tr := &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}
	// 创建一个带有自定义Transport的http.Client
	client = &http.Client{
		Transport: tr,
	}
}

func sendRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := client.Get(url)
	if err != nil {
		log.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.StatusCode)
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.example.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go sendRequest(url, &wg)
	}

	wg.Wait()

	fmt.Println("All requests completed.")
}
