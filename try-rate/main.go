package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func main() {
	r := gin.Default()
	limiter := rate.NewLimiter(rate.Limit(1), 10)
	r.Use(middleware(limiter))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.Run(":8083")
}

func middleware(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := limiter.Wait(c); err != nil {
			c.JSON(http.StatusTooManyRequests, "Forbidden request")
			c.Abort()
		}
		c.Next()
	}
}
