package handler

import (
	"net/http"
	"pln/helper"
	"pln/penggunaan"

	"github.com/gin-gonic/gin"
)

type penggunaanHandler struct {
	service penggunaan.Service
}

func NewPenggunaanHandler(service penggunaan.Service) *penggunaanHandler {
	return &penggunaanHandler{service}
}

func (h *penggunaanHandler) CreatePenggunaan(c *gin.Context) {
	var input penggunaan.PenggunaanInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Penggunaan failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPenggunaan, err := h.service.InputPenggunaan(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Penggunaan failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penggunaan.FormatterPenggunaan(newPenggunaan)
	response := helper.APIResponse("Successfuly Penggunaan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *penggunaanHandler) GetPenggunaan(c *gin.Context) {
	IdPenggunaan := c.Query("id_penggunaan")		
	penggunaans, err := h.service.GetPenggunaan(IdPenggunaan)
	if err != nil {		
		response := helper.APIResponse("Error to get penggunaan", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List of penggunaan", http.StatusOK, "success", penggunaans)
	c.JSON(http.StatusOK, response)	
}

func (h *penggunaanHandler) UpdatePenggunaan(c *gin.Context) {
	IdPenggunaan := c.Param("id_penggunaan")
	var input penggunaan.ReqUpdatePenggunaan
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update penggunaan failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatePenggunaan, err := h.service.UpdatePenggunaan(IdPenggunaan, input)
	if err != nil {
		response := helper.APIResponse("Update penggunaan failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penggunaan.FormatterPenggunaan(updatePenggunaan)
	response := helper.APIResponse("Successfuly Update Penggunaan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *penggunaanHandler) DeletePenggunaan(c *gin.Context) {
	IdPenggunaan := c.Param("id_penggunaan")
	if err := h.service.DeletePenggunaan(IdPenggunaan); err != nil {
		response := helper.APIResponse("Delete penggunaan failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}	
	response := helper.APIResponse("Successfuly Delete Penggunaan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}