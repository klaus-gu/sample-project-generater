package main

import (
	"github.com/gin-gonic/gin"
	"sample-project-generater/generater"
)

func main() {
	//engine:=gin.New()
	//engine.Use(timeoutMiddleware(time.Second * 2))
	//engine.GET("/build",timedHandler(time.Second * 15))
	r := gin.Default()
	r.GET("/build", func(c *gin.Context) {
		groupId := c.Query("groupId")
		artifactId := c.Query("artifactId")
		version := c.Query("version")
		viewName := c.Query("viewName")
		err := generater.Generate(groupId, artifactId, version, viewName)
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

//func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
//	return func(c *gin.Context) {
//		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
//
//		defer func() {
//			if ctx.Err() == context.DeadlineExceeded {
//				c.Writer.WriteHeader(http.StatusGatewayTimeout)
//				c.Abort()
//			}
//
//			//cancel to clear resources after finished
//			cancel()
//		}()
//		c.Request = c.Request.WithContext(ctx)
//		c.Next()
//	}
//}
//
//func timedHandler(duration time.Duration) func(c *gin.Context) {
//	return func(c *gin.Context) {
//		ctx := c.Request.Context()
//		type responseData struct {
//			status int
//			body   map[string]interface{}
//		}
//		doneChan := make(chan responseData)
//		go func() {
//			time.Sleep(duration)
//			doneChan <- responseData{
//				status: 200,
//				body:   gin.H{"hello": "world"},
//			}
//		}()
//		select {
//		case <-ctx.Done():
//			return
//		case res := <-doneChan:
//			c.JSON(res.status, res.body)
//		}
//	}
//}
