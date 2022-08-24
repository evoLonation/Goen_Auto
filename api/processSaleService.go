package api

import (
	"Auto/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func makeNewSale(ctx *gin.Context) {
	ret, err := service.MakeNewSale()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"bool": ret})
	return
}

type enterItemRequest struct {
	Barcode  int `json:"barcode"`
	Quantity int `json:"quantity"`
}

func enterItem(ctx *gin.Context) {
	var req enterItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//证明请求对于该结构体并不有效
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ret, err := service.EnterItem(req.Barcode, req.Quantity)
	var errPostCondition *service.ErrPostCondition
	if errors.Is(err, service.ErrPreConditionUnsatisfied) {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else if errors.As(err, &errPostCondition) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.Unwrap(err)))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"bool": ret})
	return
}
