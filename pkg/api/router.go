package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhangjie2012/go-project-layout-template/pkg/api/v1"
)

func RegisterRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")
	apiv1.Use(LoginRequired())
	RegisterRouterV1(apiv1)
}

func RegisterRouterV1(r *gin.RouterGroup) {
	r.GET("/users", v1.ListUsers)
}
