package handler

import (
	"clean-arc/model"
	"clean-arc/service"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.IService
}

type IHandler interface {
	HandleCalculator(c *gin.Context)
	GetHistoryCalculator(c *gin.Context)
}

// init factory
func NewHandler(svc service.IService) IHandler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) GetHistoryCalculator(c *gin.Context) {
	// logic khusus buat auth
	// get header username header password
	data := h.Service.GetHistoryCalculator()
	c.JSON(http.StatusOK, data)
}

// controller / handler
func (h *Handler) HandleCalculator(c *gin.Context) {
	// cuma berfokus ke kebutuhan external (request dan response)
	// logic services
	// get request body
	var reqBody model.RequestCalculator

	// binding request body
	err := c.ShouldBindJSON(&reqBody)

	// logic khusus buat auth

	// handle errror struct
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrResponse(err.Error()))
		return
	}

	allowOperator := slices.Contains([]string{"+", "-", "*", "/", "%"}, reqBody.Operator)
	if !allowOperator {
		c.JSON(http.StatusBadRequest, model.ErrResponse("invalid operator, only allow : +, -, *, /, %"))
		return
	}

	res, err := h.Service.Calculator(reqBody.A, reqBody.B, reqBody.Operator)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrResponse(err.Error()))
		return
	}

	// response json response
	var resData = model.ResponseCalculator{
		Result: res,
	}

	c.JSON(http.StatusOK, resData)
}
