package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-gin-prometheus"
)

func main() {
	r := gin.New()
	p := ginprometheus.NewPrometheus("gin")

	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}
