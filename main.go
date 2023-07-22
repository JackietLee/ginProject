package main

import (
	"fmt"
	"ginProject/routers"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")

}

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件")
	c.Next()
	//中止
	//c.Abort()
	fmt.Println("2-我是一个中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("1-我是一个中间件2")
	c.Next()
	//中止
	//c.Abort()
	fmt.Println("2-我是一个中间件2")
}

func main() {
	r := gin.Default()
	//自定义函数模板
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	// 加载模板
	//r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.Use(initMiddleware, initMiddlewareTwo)

	routers.AdminRoutersInit(r)
	routers.DefaultRoutersInit(r)
	//全局中间件
	//r.GET("/ping", func(c *gin.Context) {
	//	time.Sleep(time.Second)
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.GET("/ping", initMiddleware, initMiddlewareTwo, func(c *gin.Context) {
	//	time.Sleep(time.Second)
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
