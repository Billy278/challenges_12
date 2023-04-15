package product

import (
	"context"
	"time"

	modelProduct "github.com/Billy278/challenges_12-13/module/models/user"
)

type ProductRepo interface {
	AdmCreateProduct(ctx context.Context, productIn modelProduct.Product) (err error)
	AdmFindbyIdProduct(ctx context.Context, productId uint64) (product modelProduct.Product, err error)
	AdmFindAllProduct(ctx context.Context) (products []modelProduct.Product, err error)
	AdmUpdateProduct(ctx context.Context, productIn modelProduct.Product) (err error)
	AdmDeleteProduct(ctx context.Context, productId uint64, delete_at *time.Time) (err error)
	UsrCreateProduct(ctx context.Context, productIn modelProduct.Product) (err error)
	UsrFindAllProduct(ctx context.Context, userId int64) (products []modelProduct.Product, err error)
	UsrFindByIdProduct(ctx context.Context, productId uint64, userId uint64) (product modelProduct.Product, err error)
}
