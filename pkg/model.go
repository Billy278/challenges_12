package pkg

import (
	modelProduct "github.com/Billy278/challenges_12-13/module/models/user"
	modeluser "github.com/Billy278/challenges_12-13/module/models/user"
)

func ToProductResponse(product modelProduct.Product) modelProduct.ProductRes {
	return modelProduct.ProductRes{
		Id:         product.Id,
		UserId:     product.UserId,
		Name:       product.Name,
		Deskripsi:  product.Deskripsi,
		DetailUser: product.DetailUser,
	}
}

func ToProductsResponses(products []modelProduct.Product) (resProduct []modelProduct.ProductRes) {
	for _, product := range products {
		resProduct = append(resProduct, ToProductResponse(product))
	}
	return
}

func ToUserResponse(user modeluser.User) modeluser.UserRes {
	return modeluser.UserRes{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
		Product:  user.Product,
		Products: user.Products,
	}
}
