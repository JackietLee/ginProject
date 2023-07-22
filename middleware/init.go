package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "张三")
	context := c.Copy()
	go func() {
		fmt.Println(context.Get("username"))
		fmt.Println(context.Request.URL.Path)
	}()
}
