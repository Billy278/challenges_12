package server

import (
	"github.com/Billy278/challenges_12-13/db"
	controllerProduct "github.com/Billy278/challenges_12-13/module/controller/product"
	controllerUser "github.com/Billy278/challenges_12-13/module/controller/user"
	repositoryproduct "github.com/Billy278/challenges_12-13/module/repository/product"
	repositoryUser "github.com/Billy278/challenges_12-13/module/repository/user"
	serviceProduct "github.com/Billy278/challenges_12-13/module/service/product"
	serviceUser "github.com/Billy278/challenges_12-13/module/service/user"
)

type Ctrs struct {
	UserCtrl    controllerUser.UserCtrl
	ProductCtrl controllerProduct.ProductCtrl
}

func NewSetup() Ctrs {
	datastore := db.NewDBPostges()
	//user
	repoUser := repositoryUser.NewUserRepoImpl(datastore)
	servUser := serviceUser.NewUserSrvImpl(repoUser)
	ctrlUser := controllerUser.NewUserCtrlImpl(servUser)

	//product
	repoProduct := repositoryproduct.NewProductRepoImpl(datastore)
	servProduct := serviceProduct.NewProductSrvImpl(repoProduct)
	ctrlProduct := controllerProduct.NewProductCtrlImpl(servProduct)

	return Ctrs{
		UserCtrl:    ctrlUser,
		ProductCtrl: ctrlProduct,
	}
}
