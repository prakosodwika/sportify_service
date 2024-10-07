package handlers

import (
	"fmt"
	"sportify_services/helpers"
	"sportify_services/requests"
	"sportify_services/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type player struct {
	service services.Player
}


func NewPlayerHandler(s services.Player) *player {
	return &player{s}
}

func (h *player) CreatePlayer(c *gin.Context) {
	var r requests.Player
	response := helpers.NewResponse(c)
	if r.Latitude == "" {r.Latitude = "-"}
	if r.Longitude == "" {r.Longitude = "-"}
	
	err := c.ShouldBindJSON(&r)
	fmt.Println(&r)
	if err != nil {
		response.RequestFailed(err)
		return
	}

	if r.UserId == "" {r.UserId = r.Name}
	
	result , err := h.service.Create(r)
	if err != nil {
		response.ServiceFailed(err)
		return
	}

	response.Success(result, "Data berhasil diinputkan")
}


func (h *player) GetAllPlayers(c *gin.Context) {
	response := helpers.NewResponse(c)
	result, err := h.service.FindAll()

	if err != nil {
		response.ServiceFailed(err)
		return
	}

	response.Success(result, "Data Players")
}

func (h *player) GetPlayerById(c *gin.Context) {
	response := helpers.NewResponse(c)

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		response.RequestFailed(err)
		return
	}
	
	result, err := h.service.FindById(id)
	if err != nil {
		response.ServiceFailed(err)
		return
	}

	response.Success(result, "Data Player")
}

func (h *player) UpdatePlayer(c *gin.Context) {
	var r requests.Player
	response := helpers.NewResponse(c)

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		response.RequestFailed(err)
		return
	}

	result, err := h.service.Update(id, r)
	if err != nil {
		response.ServiceFailed(err)
		return
	}

	response.Success(result, "Data Berhasil diupdate")
}

func (h *player) DeletePlayerById(c *gin.Context) {
	response := helpers.NewResponse(c)

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		response.RequestFailed(err)
		return
	}

	result, err := h.service.Delete(id)
	if err != nil {
		response.ServiceFailed(err)
		return
	}

	response.Success(result, "Data Berhasil dihapus")
}