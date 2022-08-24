package api

import (
	"github.com/gin-gonic/gin"
)

func Start() error {

	router := gin.Default()
	processSaleServiceRouter := router.Group("/process-sale-service")
	processSaleServiceRouter.POST("/make-new-sale", makeNewSale)
	processSaleServiceRouter.POST("/enter-item", enterItem)
	//processSaleServiceRouter.POST("/end-sale", endSale)
	//processSaleServiceRouter.POST("/make-cash-payment", makeCashPayment)
	//processSaleServiceRouter.POST("/make-card-payment", makeCardPayment)

	return router.Run()
}

//该函数返回一个gin.H，gin.H是一个map，存储着键值对，将要返回给请求者
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
