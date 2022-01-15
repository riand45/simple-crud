package handler

import (
	"fmt"
	"majoo/helper"
	"majoo/repository/product"
	"majoo/repository/user"
	"majoo/services/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService 	product.Service
	authService 	auth.Service
}

func NewProductHandler(productService product.Service, authService auth.Service) *productHandler {
	return &productHandler{productService, authService}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	var input product.GetOutletProductInput
	
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get outlet's product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	input.User = currentUser
	
	products, err := h.productService.FindProducts(userID, input)
	if err != nil {
		fmt.Println(err)
		response := helper.APIResponse("Failed to get outlet's products, you are not an owner of the outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Outlet's products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var inputOutlet product.GetOutletProductInput
	
	var input product.CreateProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IDOutlet, _ := strconv.Atoi(c.Param("id"))
	input.OutletID = IDOutlet
	
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	inputOutlet.User = currentUser
	
	newProduct, err := h.productService.CreateProduct(userID, inputOutlet, input)
	if err != nil {
		response := helper.APIResponse("Failed to get outlet's product", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create product", http.StatusOK, "success", newProduct)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UploadImage(c *gin.Context) {
	var input product.CreateProductImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload product image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload product image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload product image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.SaveProductImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload product image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Product image successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}