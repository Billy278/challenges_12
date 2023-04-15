package product

import (
	"context"

	modelProduct "github.com/Billy278/challenges_12-13/module/models/user"
)

type ProductSrv interface {
	AdmCreateProductSrv(ctx context.Context, productIn modelProduct.Product) (err error)
	AdmFindbyIdProuctSrv(ctx context.Context, productId uint64) (productRes modelProduct.ProductRes, err error)
	AdmFindAllProductSrv(ctx context.Context) (productsRes []modelProduct.ProductRes, err error)
	AdmUpdateProductSrv(ctx context.Context, productIn modelProduct.Product) (err error)
	AdmDeleteProductSrv(ctx context.Context, productId uint64) (err error)
	UsrCreateProductSrv(ctx context.Context, productIn modelProduct.Product) (err error)
	UsrFindAllProductSrv(ctx context.Context, userId int64) (productsRes []modelProduct.ProductRes, err error)
	UsrFindByIdProductSrv(ctx context.Context, productId uint64, userId uint64) (product modelProduct.ProductRes, err error)
}
