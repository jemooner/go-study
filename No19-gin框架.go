/*
这个示例代码中创建了一个Gin默认的路由引擎`router`，并通过`router.GET`和`router.POST`定义了两个路由规则，分别处理GET请求和POST请求。
在GET请求的处理函数里，返回一个JSON响应，而在POST请求的处理函数里，获取表单参数，并以JSON响应的形式返回。

最后调用`router.Run(":8080")`启动HTTP服务器，并监听在8080端口。
你可以通过访问`http://localhost:8080`来测试GET请求，以及使用工具（如Postman）发送POST请求到`http://localhost:8080/login`来测试POST请求。
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个Gin默认的路由引擎
	router := gin.Default()

	// 定义一个GET请求的路由规则
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// 定义一个POST请求的路由规则
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// 启动HTTP服务器，并监听在8080端口
	router.Run(":8080")
}
