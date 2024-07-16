package handler

import (
	"net/http"
	"pln/helper"
	"pln/level"

	"github.com/gin-gonic/gin"
)



type levelHandler struct {
	service level.Service
}

func NewLevelHandler(service level.Service) *levelHandler {
	return &levelHandler{service}
}

func (h *levelHandler) CreateLevel(c *gin.Context) {
	var input level.LevelInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("level failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newLevel, err := h.service.CreateLevel(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("level failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := level.FormatterLevel(newLevel)
	response := helper.APIResponse("Successfuly level", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}


func (h *levelHandler) GetLevel(c *gin.Context) {
	level, err := h.service.GetLevel()

	if err != nil {		
		response := helper.APIResponse("Error to get level", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List of level", http.StatusOK, "success", level)
	c.JSON(http.StatusOK, response)	
}