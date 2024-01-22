package main

import (
	"fmt"
	"gin-example/internal/b64"
	"gin-example/internal/httpmiddleware"
	"gin-example/internal/unixtime"
	"github.com/gin-gonic/gin"
	"net/http"
)

// curl -v localhost:8080/unixtime/3523463425

func main() {
	router := gin.Default()

	router.Use(httpmiddleware.SetRequestIdHeader())
	router.Use(httpmiddleware.LogRequest())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	b64.RegisterRoutes(router)
	unixtime.RegisterRoutes(router)

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("http://localhost:8080")
	srv.ListenAndServe()
}
