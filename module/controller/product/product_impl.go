package product

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Billy278/challenges_12-13/module/models/token"
	model "github.com/Billy278/challenges_12-13/module/models/user"
	service "github.com/Billy278/challenges_12-13/module/service/product"
	response "github.com/Billy278/challenges_12-13/pkg"
	helper "github.com/Billy278/challenges_12-13/pkg/cripto"
	"github.com/Billy278/challenges_12-13/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type ProductCtrlImpl struct {
	ProductService service.ProductSrv
}

func NewProductCtrlImpl(productservice service.ProductSrv) ProductCtrl {
	return &ProductCtrlImpl{
		ProductService: productservice,
	}
}

func (p_ctrl *ProductCtrlImpl) CreateProductCtrl(ctx *gin.Context) {
	reqIn := model.Product{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	//ingat nnti userid sama rolenya kita input manual
	if reqIn.Name == "" || reqIn.Deskripsi == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "failed to Insert Product",
		})
		return
	}
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid user id",
		})
		return
	}
	fmt.Println("isi accessclaims sebelum di parse ke mapping")
	fmt.Println(accessClaimIn)
	fmt.Println("=========O================================")
	var accessClaim token.AccessClaim
	if err := helper.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid payload",
		})
		return
	}

	Role := accessClaim.Role
	userId, _ := strconv.Atoi(accessClaim.UserID)
	reqIn.UserId = int64(userId)
	if Role == "admin" {
		err := p_ctrl.ProductService.AdmCreateProductSrv(ctx, reqIn)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
				Code:   http.StatusBadRequest,
				Status: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "Success Create Product ",
		})
	} else {
		err := p_ctrl.ProductService.UsrCreateProductSrv(ctx, reqIn)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
				Code:   http.StatusBadRequest,
				Status: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "Success Create Product ",
		})
	}
}
func (p_ctrl *ProductCtrlImpl) FindbyIdProuctCtrl(ctx *gin.Context) {
	id, err := p_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	//=======================================================================
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid user id",
		})
		return
	}
	var accessClaim token.AccessClaim
	if err := helper.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid payload",
		})
		return
	}

	Role := accessClaim.Role
	userIdctx, _ := strconv.Atoi(accessClaim.UserID)
	user_id := userIdctx
	//=======================================================================

	if Role == "admin" {
		product, err := p_ctrl.ProductService.AdmFindbyIdProuctSrv(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.WebResponse{
				Code:   http.StatusNotFound,
				Status: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "Success Find by id Product",
			Data:   product,
		})
	} else {
		product, err := p_ctrl.ProductService.UsrFindByIdProductSrv(ctx, id, uint64(user_id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.WebResponse{
				Code:   http.StatusNotFound,
				Status: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "Success Find by id Product",
			Data:   product,
		})
	}
}
func (p_ctrl *ProductCtrlImpl) FindAllProductCtrl(ctx *gin.Context) {
	//=======================================================================
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid user id",
		})
		return
	}
	var accessClaim token.AccessClaim
	if err := helper.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid payload",
		})
		return
	}

	Role := accessClaim.Role
	userIdctx, _ := strconv.Atoi(accessClaim.UserID)
	user_id := userIdctx
	//=======================================================================
	if Role == "admin" {
		product, err := p_ctrl.ProductService.AdmFindAllProductSrv(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.WebResponse{
				Code:   http.StatusNotFound,
				Status: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "Success Find All Product",
			Data:   product,
		})
	} else {
		product, err := p_ctrl.ProductService.UsrFindAllProductSrv(ctx, int64(user_id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.WebResponse{
				Code:   http.StatusNotFound,
				Status: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.WebResponse{
			Code:   http.StatusOK,
			Status: "Success Find All Product",
			Data:   product,
		})
	}
}
func (p_ctrl *ProductCtrlImpl) UpdateProductCtrl(ctx *gin.Context) {
	//=======================================================================
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid user id",
		})
		return
	}
	var accessClaim token.AccessClaim
	if err := helper.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid payload",
		})
		return
	}

	Role := accessClaim.Role

	//=======================================================================
	if Role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "NOT Unauthorized ",
		})
		return
	}
	id, err := p_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	reqIn := model.Product{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	if reqIn.Name == "" || reqIn.Deskripsi == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "failed to Update Product",
		})
		return
	}
	reqIn.Id = int64(id)
	err = p_ctrl.ProductService.AdmUpdateProductSrv(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Update Product",
	})
}
func (p_ctrl *ProductCtrlImpl) DeleteProductCtrl(ctx *gin.Context) {
	//=======================================================================
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid user id",
		})
		return
	}
	var accessClaim token.AccessClaim
	if err := helper.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "invalid payload",
		})
		return
	}

	Role := accessClaim.Role

	//=======================================================================
	if Role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "NOT Unauthorized ",
		})
		return
	}
	id, err := p_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	err = p_ctrl.ProductService.AdmDeleteProductSrv(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Delete Product",
	})
}

func (p_ctrl *ProductCtrlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")
	if id == "" {
		err = errors.New("failed id")
		ctx.JSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		ctx.JSON(http.StatusBadRequest, response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	return idUint, err

}
