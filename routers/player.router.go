package routers

import (
	"sportify_services/config"
	"sportify_services/handlers"
	"sportify_services/repositories"
	"sportify_services/services"

	"github.com/gin-gonic/gin"
)

func player(r *gin.RouterGroup) {
	db := config.DB

	repository := repositories.NewPlayerRepository(db)
	service    := services.NewPlayerService(repository)
	handler    := handlers.NewPlayerHandler(service)

	r.POST("/create", handler.CreatePlayer)
	r.GET("/getAll", handler.GetAllPlayers)
	r.GET("/getById/:id", handler.GetPlayerById)
	r.PUT("/update/:id", handler.UpdatePlayer)
	r.DELETE("/delete/:id", handler.DeletePlayerById)
}