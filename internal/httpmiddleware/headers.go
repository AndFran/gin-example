package httpmiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetRequestIdHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		guid := uuid.New().String()
		c.Set("uuid", guid)
		c.Header("X-Request-id", guid)
	}
}

