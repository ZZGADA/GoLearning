package ginstudy

import (
	"basicLearning/src/entity"
	"fmt"
	"github.com/gin-gonic/gin"
)

func webDemo() *gin.Engine {
	r := gin.Default()
	return r
}

func webDefault() {
	r := webDemo()
	r.GET("/goHello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}

func WebRouter() {
	router := gin.Default()

	/*
		通过":id"来指定参数名，在处理器函数中使用c.Param(“id”)获取参数值，
		最后返回给客户端。我们也使用Query方法获取查询参数，使用"c.Query(“query”)"获取名为"query"的参数值，最后返回给客户端。
	*/

	// GET 请求处理
	router.GET("/hello", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// /book/123  动态路由
	router.GET("/book/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "Book ID is "+id)
	})

	// /search?query=gin  路由携带参数
	router.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		more := c.Query("more")
		c.String(200, "Search query is "+query+" more is "+more)
	})

	// POST 请求处理
	router.POST("/users", func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 处理接收到的用户数据
		// ...
		fmt.Println(user.Name)
		fmt.Println(user)
		c.JSON(200, gin.H{
			"message": "User created successfully",
		})
	})

	router.Run(":8080")
}
