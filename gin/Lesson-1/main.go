package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		name := context.Query("name")
		context.String(http.StatusOK, name)
		fmt.Println(name)
	})

	r.Run(":8083")
}
