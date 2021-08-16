package main

import ("github.com/gin-gonic/gin"
_"net/http")


	

	func main() {
		router := gin.Default()
		router.POST("/article", func(c *gin.Context) {
			c.String(200, "article post")
		})
		router.DELETE("/article", func(c *gin.Context) {
			c.String(200, "article delete")
		})
		router.PUT("/article", func(c *gin.Context) {
			c.String(200, "article put")
		})
		router.GET("/article", func(c *gin.Context) {
			c.String(200, "article get")
		})
		router.Run(":3033")
	}
