package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"koratta":  "koratta",
		"koratta2": "koratta2",
	}))
	authorized.GET("/hello", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(200, gin.H{"message": "Hello " + user})
	})
	r.Run(":8080")
}
