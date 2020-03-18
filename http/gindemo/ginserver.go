package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

const keyId = "requestId"

func main() {
	//打印版本号
	//fmt.Println(runtime.Version())
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	//gin的中间件，在这里可以执行我们想要干的事
	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()
		logger.Info("incoming request:", zap.String("path", c.Request.URL.Path), zap.Int("status", c.Writer.Status()), zap.Duration("time duration", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		c.Set(keyId, time.Now())
	})

	r.GET("/ping", func(c *gin.Context) {
		requestId, b := c.Get(keyId)
		h := gin.H{
			"message": "pong",
		}
		if b {
			h[keyId] = requestId
		}
		c.JSON(200, h)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
