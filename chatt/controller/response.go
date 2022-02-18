package controller

import "github.com/gin-gonic/gin"

func ErrorResponseJson(c *gin.Context, code int, message string, err error) {
	c.JSON(code, gin.H{
		"detail":  err.Error(),
		"message": message,
	})
	return
}

func SuccessResponseJson(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"message": message,
		"data":    data,
	})
	return
}
