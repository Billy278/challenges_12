package user

import (
	"net/http"

	model "github.com/Billy278/challenges_12-13/module/models/user"
	service "github.com/Billy278/challenges_12-13/module/service/user"
	response "github.com/Billy278/challenges_12-13/pkg"
	"github.com/gin-gonic/gin"
)

type UserCtrlImpl struct {
	UserSrvice service.UserSrv
}

func NewUserCtrlImpl(userservice service.UserSrv) UserCtrl {
	return &UserCtrlImpl{
		UserSrvice: userservice,
	}
}

func (u_ctrl *UserCtrlImpl) Registration(ctx *gin.Context) {
	reqIn := model.User{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	if reqIn.Name == "" || reqIn.Username == "" || reqIn.Password == "" || reqIn.Role == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "failed to Registration",
		})
		return
	}

	_, err := u_ctrl.UserSrvice.FindbyUsername(ctx, reqIn.Username)
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "USERNAME SUDAH ADA",
		})
		return
	}
	err = u_ctrl.UserSrvice.CreateUser(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Registration ",
	})
}

func (u_ctrl *UserCtrlImpl) LoginUser(ctx *gin.Context) {
	reqIn := model.User{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	if reqIn.Username == "" || reqIn.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "failed to login",
		})
		return
	}
	tokens, err := u_ctrl.UserSrvice.LoginByUserName(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "suscress Login And Get Tokens",
		Data:   tokens,
	})

}
