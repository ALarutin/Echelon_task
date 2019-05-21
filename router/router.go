package router

import (
	"github.com/ALarutin/Echelon_task/handlers"
	"github.com/gin-gonic/gin"
)

func GetRouter() (r *gin.Engine) {
	r = gin.Default()
	r.POST("/checkText", handlers.CheckText)
	return
}
