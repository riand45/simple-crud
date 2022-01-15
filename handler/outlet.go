package handler

import (
	"fmt"
	"majoo/helper"
	"majoo/repository/outlet"
	"majoo/repository/user"
	"majoo/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type outletHandler struct {
	outletService 	outlet.Service
	authService 	auth.Service
}

func NewOutletHandler(outletService outlet.Service, authService auth.Service) *outletHandler {
	return &outletHandler{outletService, authService}
}

func (h *outletHandler) GetOutlets(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	outlets, err := h.outletService.FindOutlets(userID)
	if err != nil {
		response := helper.APIResponse("Error to get outlets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of outlets", http.StatusOK, "success", outlets)
	c.JSON(http.StatusOK, response)
}

func (h *outletHandler) GetOutlet(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	fmt.Println("userid:",currentUser)
	var input outlet.GetOutletDetailInput
	
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	outletDetail, err := h.outletService.FindOutlet(userID, input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Outlet detail", http.StatusOK, "success", outletDetail)
	c.JSON(http.StatusOK, response)
}

func (h *outletHandler) CreateOutlet(c *gin.Context) {
	var input outlet.CreateOutletInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create outlet", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	fmt.Println(currentUser)
	input.User = currentUser

	newOutlet, err := h.outletService.CreateOutlet(input)
	if err != nil {
		response := helper.APIResponse("Failed to create outlet", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create outlet", http.StatusOK, "success", newOutlet)
	c.JSON(http.StatusOK, response)
}

func (h *outletHandler) UpdateOutlet(c *gin.Context) {
	var inputID outlet.GetOutletDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData outlet.CreateOutletInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {	
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update outlet", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	fmt.Println("userid:",userID)
	inputData.User = currentUser

	updateOutlet, err := h.outletService.UpdateOutlet(userID, inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Not an owner of the outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	respnse := helper.APIResponse("Success to update outlet", http.StatusOK, "success", updateOutlet)
	c.JSON(http.StatusOK, respnse)
}

func (h *outletHandler) DeleteOutlet(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	
	var input outlet.GetOutletDetailInput
	
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get delete of outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.outletService.DeleteOutlet(userID, input)
	if err != nil {
		response := helper.APIResponse("Failed to get delete of outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Outlet delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}