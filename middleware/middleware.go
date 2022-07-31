package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is custom middleware")
		fmt.Printf("request method : %s \nrequest host: %s\n ",
			c.Request.Method, c.Request.Host)
		c.Next()
	}
}
