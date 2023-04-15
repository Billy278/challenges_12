package server

import (
	"github.com/Billy278/challenges_12-13/module/router"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	ctrl := NewSetup()
	rt := gin.Default()
	//init middleware

	rt.Use(gin.Recovery(), gin.Logger())
	router.NewRouter(rt, ctrl.UserCtrl, ctrl.ProductCtrl)
	rt.Run("0.0.0.0:8080")
}
