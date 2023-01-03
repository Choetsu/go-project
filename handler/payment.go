package handler

import (
	"io"
	"net/http"
	"strconv"

	"go-project/payment"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.Service
}

func NewPaymenthandler(paymentService payment.Service) *paymentHandler {
	return &paymentHandler{paymentService}
}

func (h *paymentHandler) GetAll(c *gin.Context) {
	payments, err := h.paymentService.GetAll()
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
		Data:    payments,
	})
}

func (h *paymentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "ID invalide",
			Data:    err.Error(),
		})
		return
	}

	payment, err := h.paymentService.GetByID(id)
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
		Message: "OK",
		Data:    payment,
	})
}

func (h *paymentHandler) Create(c *gin.Context) {
	var input payment.InputPayment
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		})
		return
	}

	newPayment, err := h.paymentService.Create(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "OK",
		Data:    newPayment,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *paymentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "ID invalide",
			Data:    err.Error(),
		})
		return
	}

	var input payment.InputPayment
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uPayment, err := h.paymentService.Update(id, input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Une erreur est survenue",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "OK",
		Data:    uPayment,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *paymentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "ID invalide",
			Data:    err.Error(),
		})
		return
	}

	err = h.paymentService.Delete(id)
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
		Message: "OK",
	})
}

func (h *paymentHandler) StreamAll(c *gin.Context) {
	streams, err := h.paymentService.StreamAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    nil,
		})
		return
	}

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", streams)
		return true
	})
}
