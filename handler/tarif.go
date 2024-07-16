package handler

import (
	"net/http"
	"pln/helper"
	"pln/tarif"

	"github.com/gin-gonic/gin"
)

type tarifHandler struct {
	service tarif.Service
}

func NewTarifHandler(service tarif.Service) *tarifHandler {
	return &tarifHandler{service}
}

func (h *tarifHandler) CreateTarif(c *gin.Context) {
	var input tarif.TarifInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Tarif failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTarif, err := h.service.TarifInputUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Tarif failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := tarif.TarifFormatter(newTarif)
	response := helper.APIResponse("Successfuly Tarif", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *tarifHandler) GetTarif(c *gin.Context) {
	tarif, err := h.service.GetTarif()

	if err != nil {		
		response := helper.APIResponse("Error to get tarif", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List of tarif", http.StatusOK, "success", tarif)
	c.JSON(http.StatusOK, response)	
}

func (h *tarifHandler) UpdateTarif(c *gin.Context) {
	idTarif := c.Param("id_tarif")
	var input tarif.ReqUpdateTarif
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update tarif failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateTarif, err := h.service.UpdateTarif(idTarif, input)
	if err != nil {
		response := helper.APIResponse("Update tarif failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := tarif.FormatterTarif(updateTarif)
	response := helper.APIResponse("Successfuly Update tarif", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *tarifHandler) DeleteTarif(c *gin.Context) {
	IdTarif := c.Param("id_tarif")
	if err := h.service.DeleteTarif(IdTarif); err != nil {
		response := helper.APIResponse("Delete tarif failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}	
	response := helper.APIResponse("Successfuly Delete tarif", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}