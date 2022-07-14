package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "update app baru dicoba 1",
			"data":    []string{},
		})
	})
	r.Run(":8881")
}
