package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
