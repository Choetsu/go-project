package handler

import (
	"net/http"
	"strconv"

	"go-project/product"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetAll(c *gin.Context) {
	products, err := h.productService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "OK",
		Data:    products,
	})
}

func (h *productHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "ID invalide",
			Data:    err.Error(),
		})
		return
	}

	product, err := h.productService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "OK",
		Data:    product,
	})
}

func (h *productHandler) Create(c *gin.Context) {
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Données invalides",
			Data:    err.Error(),
		})
		return
	}

	newProduct, err := h.productService.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		})
		return
	}
	response := &Response{
		Success: true,
		Message: "Produit créé",
		Data:    newProduct,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *productHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "ID invalide",
			Data:    err.Error(),
		})
		return
	}

	var input product.InputProduct
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Données invalides",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uProduct, err := h.productService.Update(id, input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
	}
	response := &Response{
		Success: true,
		Message: "Produit mis à jour",
		Data:    uProduct,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *productHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "ID invalide",
			Data:    err.Error(),
		})
		return
	}

	err = h.productService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Produit supprimé",
	})
}
