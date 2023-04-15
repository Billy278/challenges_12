package product

import (
	"context"
	"log"
	"time"

	modelProduct "github.com/Billy278/challenges_12-13/module/models/user"
	repository "github.com/Billy278/challenges_12-13/module/repository/product"
	helper "github.com/Billy278/challenges_12-13/pkg"
)

type ProductSrvImpl struct {
	ProductRepo repository.ProductRepo
}

func NewProductSrvImpl(productrepo repository.ProductRepo) ProductSrv {
	return &ProductSrvImpl{
		ProductRepo: productrepo,
	}
}

func (p_srv *ProductSrvImpl) AdmCreateProductSrv(ctx context.Context, productIn modelProduct.Product) (err error) {
	tNow := time.Now()
	productIn.Create_at = &tNow
	err = p_srv.ProductRepo.AdmCreateProduct(ctx, productIn)
	if err != nil {
		log.Printf("[ERROR] ADMIN error  create Product :%v\n", err)
		return
	}
	return
}
func (p_srv *ProductSrvImpl) AdmFindbyIdProuctSrv(ctx context.Context, productId uint64) (productRes modelProduct.ProductRes, err error) {
	product, err := p_srv.ProductRepo.AdmFindbyIdProduct(ctx, productId)
	if err != nil {
		log.Printf("[INFO] ID Product:%v\n", err)
		return
	}
	return helper.ToProductResponse(product), err
}
func (p_srv *ProductSrvImpl) AdmFindAllProductSrv(ctx context.Context) (productsRes []modelProduct.ProductRes, err error) {
	product, err := p_srv.ProductRepo.AdmFindAllProduct(ctx)
	if err != nil {
		log.Printf("[INFO] ID Product:%v\n", err)
		return
	}
	return helper.ToProductsResponses(product), err

}
func (p_srv *ProductSrvImpl) AdmUpdateProductSrv(ctx context.Context, productIn modelProduct.Product) (err error) {
	product, err := p_srv.ProductRepo.AdmFindbyIdProduct(ctx, uint64(productIn.Id))
	if err != nil {
		log.Printf("[INFO] ID Product:%v\n", err)
		return
	}
	product.Name = productIn.Name
	product.Deskripsi = productIn.Deskripsi
	tNow := time.Now()
	product.Update_at = &tNow
	err = p_srv.ProductRepo.AdmUpdateProduct(ctx, product)
	if err != nil {
		log.Printf("[EROR] admin Update Product:%v\n", err)
		return
	}
	return

}
func (p_srv *ProductSrvImpl) AdmDeleteProductSrv(ctx context.Context, productId uint64) (err error) {
	_, err = p_srv.ProductRepo.AdmFindbyIdProduct(ctx, uint64(productId))
	if err != nil {
		log.Printf("[INFO] ID Product:%v\n", err)
		return
	}
	tNow := time.Now()
	delete_at := &tNow
	err = p_srv.ProductRepo.AdmDeleteProduct(ctx, uint64(productId), delete_at)
	if err != nil {
		log.Printf("[EROR] admin Error Product:%v\n", err)
		return
	}
	return
}
func (p_srv *ProductSrvImpl) UsrCreateProductSrv(ctx context.Context, productIn modelProduct.Product) (err error) {
	tNow := time.Now()
	productIn.Create_at = &tNow
	err = p_srv.ProductRepo.UsrCreateProduct(ctx, productIn)
	if err != nil {
		log.Printf("[ERROR] user error  create Product :%v\n", err)
		return
	}
	return
}
func (p_srv *ProductSrvImpl) UsrFindAllProductSrv(ctx context.Context, userId int64) (productsRes []modelProduct.ProductRes, err error) {
	product, err := p_srv.ProductRepo.UsrFindAllProduct(ctx, userId)
	if err != nil {
		log.Printf("[INFO] ID Product:%v\n", err)
		return
	}
	return helper.ToProductsResponses(product), err
}

func (p_srv *ProductSrvImpl) UsrFindByIdProductSrv(ctx context.Context, productId uint64, userId uint64) (productRes modelProduct.ProductRes, err error) {
	product, err := p_srv.ProductRepo.UsrFindByIdProduct(ctx, productId, userId)
	if err != nil {
		log.Printf("[INFO] ID Product:%v\n", err)
		return
	}

	return helper.ToProductResponse(product), err
}
