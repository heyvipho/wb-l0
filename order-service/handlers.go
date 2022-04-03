package main

import "github.com/gin-gonic/gin"

func getOrderHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": c.Param("id"),
	})
}
