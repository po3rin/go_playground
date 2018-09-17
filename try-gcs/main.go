package main

import (
	"go-playgroung/try-gcs/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Cache-Control", "X-Requested-With"},
	}))
	r.POST("/upload", handler.Handler)
	r.Run(":5000")
}
