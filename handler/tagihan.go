package handler

import (
	"net/http"
	"pln/helper"
	"pln/tagihan"

	"github.com/gin-gonic/gin"
)

type tagihanHandler struct {
	service tagihan.Service
}

func NewTagihanHandler(service tagihan.Service) *tagihanHandler {
	return &tagihanHandler{service}
}

func (h *tagihanHandler) CreateTagihan(c *gin.Context) {
	var input tagihan.TagihanInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Tagihan failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTagihan, err := h.service.InputService(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Tagihan failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := tagihan.FormatterTagihan(newTagihan)
	response := helper.APIResponse("Successfuly Tagihan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *tagihanHandler) GetTagihan(c *gin.Context) {
	tagihan, err := h.service.GetTagihan()

	if err != nil {		
		response := helper.APIResponse("Error to get tagihan", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List of tagihan", http.StatusOK, "success", tagihan)
	c.JSON(http.StatusOK, response)	
}

func (h *tagihanHandler) UpdateTagihan(c *gin.Context) {
	IdTagihan := c.Param("id_tagihan")
	var input tagihan.ReqUpdateTagihan
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update tarif failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateTagihan, err := h.service.UpdateTagihan(IdTagihan, input)
	if err != nil {
		response := helper.APIResponse("Update tarif failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := tagihan.FormatterTagihan(updateTagihan)
	response := helper.APIResponse("Successfuly Update tarif", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}


func (h *tagihanHandler) DeleteTagihan(c *gin.Context) {
	IdTagihan := c.Param("id_tagihan")
	if err := h.service.DeleteTagihan(IdTagihan); err != nil {
		response := helper.APIResponse("Delete tagihan failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}	
	response := helper.APIResponse("Successfuly Delete tagihan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}