package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangjie2012/go-project-layout-template/pkg/api"
)

func (s *api.Server) ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
