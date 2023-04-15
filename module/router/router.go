package router

import (
	controllerProduct "github.com/Billy278/challenges_12-13/module/controller/product"
	controllerUser "github.com/Billy278/challenges_12-13/module/controller/user"
	"github.com/Billy278/challenges_12-13/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, UserCtrl controllerUser.UserCtrl, ProductCtrl controllerProduct.ProductCtrl) {
	//register all router

	r.POST("/register", UserCtrl.Registration)
	r.POST("/login", UserCtrl.LoginUser)
	group := r.Group("/user", middleware.BearerOAuth())
	group.GET("/product", ProductCtrl.FindAllProductCtrl)
	group.GET("/product/:id", ProductCtrl.FindbyIdProuctCtrl)
	group.POST("/product", ProductCtrl.CreateProductCtrl)
	group.PUT("/product/:id", ProductCtrl.UpdateProductCtrl)
	group.DELETE("/product/:id", ProductCtrl.DeleteProductCtrl)

}
