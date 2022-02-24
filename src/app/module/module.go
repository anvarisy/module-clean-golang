package module

import "github.com/gin-gonic/gin"

type ModuleInterface interface {
	GetHttpRouter() *gin.Engine
}
