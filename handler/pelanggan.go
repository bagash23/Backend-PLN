package handler

import (
	"net/http"
	"pln/helper"
	"pln/pelanggan"

	"github.com/gin-gonic/gin"
)


type pelangganHandler struct {
	service pelanggan.Service
}

func NewPelangganHandler(service pelanggan.Service) *pelangganHandler {
	return &pelangganHandler{service}
}


func (h *pelangganHandler) CreatePelanggan(c *gin.Context) {
	var input pelanggan.PelangganInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Penggunaan failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPelanggan, err := h.service.InputPelanggan(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Penggunaan failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pelanggan.FormatterPelanggan(newPelanggan)
	response := helper.APIResponse("Successfuly Penggunaan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}


func (h *pelangganHandler) GetPelanggan(c *gin.Context) {
	tagihan, err := h.service.GetPelanggan()

	if err != nil {		
		response := helper.APIResponse("Error to get pelanggan", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List of pelanggan", http.StatusOK, "success", tagihan)
	c.JSON(http.StatusOK, response)	
}