package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// do some login check
		c.Next()
	}
}
