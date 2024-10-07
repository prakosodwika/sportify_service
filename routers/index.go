package routers

import "github.com/gin-gonic/gin"

func Index(r *gin.RouterGroup) {
	player(r.Group("/player"))
}