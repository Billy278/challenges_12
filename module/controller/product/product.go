package product

import "github.com/gin-gonic/gin"

type ProductCtrl interface {
	CreateProductCtrl(ctx *gin.Context)
	FindbyIdProuctCtrl(ctx *gin.Context)
	FindAllProductCtrl(ctx *gin.Context)
	UpdateProductCtrl(ctx *gin.Context)
	DeleteProductCtrl(ctx *gin.Context)
}
