package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("sample-project-generater.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/build", func(c *gin.Context) {
		groupId := c.Query("groupId")
		artifactId := c.Query("artifactId")
		version := c.Query("version")
		viewName := c.Query("viewName")
		err := Generate(groupId, artifactId, version, viewName)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    0,
				"result":  err.Error(),
				"message": "build failed!",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    1,
			"result":  "",
			"message": "build success!",
		})

	})
	r.Run(":7001")
}
