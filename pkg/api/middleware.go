package api

import (
	"github.com/gin-gonic/gin"
)

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// do some login check
		c.Next()
	}
}
