package handler

import (
	"net/http"
	"pln/helper"
	"pln/pembayaran"

	"github.com/gin-gonic/gin"
)

type pembayaranHandler struct {
	service pembayaran.Service
}

func NewPembayaranHandler(service pembayaran.Service) *pembayaranHandler {
	return &pembayaranHandler{service}
}

func (h *pembayaranHandler) CreatePembayaran(c *gin.Context) {
	var input pembayaran.PembayaranInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		// errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("pembayaran failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPembayaran, err := h.service.InputPembayaran(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("pembayaran failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pembayaran.FormatterPembayaran(newPembayaran)
	response := helper.APIResponse("Successfuly pembayaran", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pembayaranHandler) GetPembayaran(c *gin.Context) {
	IdPembayaran := c.Query("id_pembayaran")		
	pembayarans, err := h.service.GetPembayaran(IdPembayaran)
	if err != nil {		
		response := helper.APIResponse("Error to get pembayaran", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List of pembayaran", http.StatusOK, "success", pembayarans)
	c.JSON(http.StatusOK, response)	
}