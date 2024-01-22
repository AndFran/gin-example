package unixtime

import "github.com/gin-gonic/gin"

// curl localhost:3000/unixtime/3523463425

func RegisterRoutes(r *gin.Engine) {
	var h handler
	routesUnixtime := r.Group("/unixtime")
	routesUnixtime.GET("/:timestamp", h.toHuman)
}

