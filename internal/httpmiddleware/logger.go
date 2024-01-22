package httpmiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		guid, _ := c.Get("uuid")
		fmt.Printf("The request with uuid %s started\n", guid.(string))
		c.Next() // next in chain
		fmt.Printf("The request with uuid %s finished\n", guid.(string))
	}
}

