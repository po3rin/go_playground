package main

import (
	"go-playground/try-logrus-stackdriver/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/all", handler.AllHandler)
	r.GET("/info", handler.InfoHandler)
	r.GET("/warn", handler.WarnHandler)
	r.GET("/error", handler.ErrorHandler)
	r.Run()
}
