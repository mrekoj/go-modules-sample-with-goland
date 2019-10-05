package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": quote.Hello(),
		})
	})

	r.GET("/concurrency", func(context *gin.Context) {
		context.String(http.StatusOK, quoteV3.Concurrency())
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
