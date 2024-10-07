package helpers

import (
	"fmt"
	"net/http"
	"sportify_services/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response interface {
	RequestFailed(err error)
	ServiceFailed(err error)
	Success(data struct{}, message string)
}

type response struct {
	context *gin.Context
}

func NewResponse(c *gin.Context) *response {
	return &response{c}
}

func (r *response) RequestFailed(err error) {
	errorMessages :=  make(map[string]string)
	errorMessage := "Request Error"

	if err.Error() == "EOF" {
		errorMessage = "Request Tidak Boleh Kosong"
	} else {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages[e.Field()] = "Format salah"
		}
	}


	result := models.Failed{
		Status: "Failed",
		Code: http.StatusBadRequest,
		Message: errorMessage,
		Error: errorMessages,
	}

	r.context.JSON(http.StatusBadRequest, result)
}

func (r *response) ServiceFailed(err error) {

	var errorMessage string
	var statusCode int

	switch err.Error() {
	case "404": 
		statusCode = http.StatusNotFound
		errorMessage = "Data Tidak Ditemukan"
	default: 
		fmt.Println(err)
		statusCode = http.StatusInternalServerError
		errorMessage = "Internal Server Error"
	}

	result := models.Failed{
		Status: "Failed",
		Code: statusCode,
		Message: errorMessage,
	}

	r.context.JSON(statusCode, result)
}

func (r *response) Success(data interface{}, message string) {
	result := models.Success{
		Status: "Success",
		Code: http.StatusOK,
		Message: message,
		Data: data,
	}

	r.context.JSON(http.StatusOK, result)
}