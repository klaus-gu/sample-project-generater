package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/build", func(c *gin.Context) {
		groupId := c.Query("groupId")
		artifactId := c.Query("artifactId")
		version := c.Query("version")
		viewName := c.Query("viewName")
		err := Generate(groupId, artifactId, version, viewName)
		if err != nil {
			c.JSON(200, gin.H{
				"code":0,
				"result": err,
				"message":"build failed!",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":1,
			"result": err,
			"message":"build success!",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
