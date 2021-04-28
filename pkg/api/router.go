package api

import "github.com/gin-gonic/gin"

func (s *Server) RegisterRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")
	apiv1.Use(s.LoginRequired())
}

func (s *Server) RegisterRouterV1(v1 *gin.RouterGroup) {

}
