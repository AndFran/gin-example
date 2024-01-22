package b64

import "github.com/gin-gonic/gin"

// curl "127.0.0.1:3000/base64/encode?text=hello"
// curl "127.0.0.1:3000/base64/decode?base64=aGVsbG8="

func RegisterRoutes(r *gin.Engine) {
	h := &handler{}
	routesBase64 := r.Group("/base64")
	routesBase64.GET("/decode", h.decode)
	routesBase64.GET("/encode", h.encode)
}

