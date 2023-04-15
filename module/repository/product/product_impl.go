package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	modelProduct "github.com/Billy278/challenges_12-13/module/models/user"
)

type ProductRepoImpl struct {
	DB *sql.DB
}

func NewProductRepoImpl(db *sql.DB) ProductRepo {
	return &ProductRepoImpl{
		DB: db,
	}
}

func (p_repo *ProductRepoImpl) AdmCreateProduct(ctx context.Context, productIn modelProduct.Product) (err error) {
	logCtx := fmt.Sprintf("%T - AdmCreateProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "INSERT INTO data_product(user_id,name,deskripsi,create_at) VALUES($1,$2,$3,$4)"
	_, err = p_repo.DB.ExecContext(ctx, sql, productIn.UserId, productIn.Name, productIn.Deskripsi, productIn.Create_at)
	if err != nil {
		return
	}
	return
}
func (p_repo *ProductRepoImpl) AdmFindbyIdProduct(ctx context.Context, productId uint64) (product modelProduct.Product, err error) {
	logCtx := fmt.Sprintf("%T - AdmFindProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,user_id,name,deskripsi FROM data_product WHERE id=$1 AND delete_at IS NULL"
	rows, err := p_repo.DB.QueryContext(ctx, sql, productId)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&product.Id, &product.UserId, &product.Name, &product.Deskripsi)
		if err != nil {
			return
		}
		return
	} else {
		return product, errors.New("PRODUCT NOT FOUND")
	}
}
func (p_repo *ProductRepoImpl) AdmFindAllProduct(ctx context.Context) (products []modelProduct.Product, err error) {
	logCtx := fmt.Sprintf("%T - AdmFindAllProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,user_id,name,deskripsi FROM data_product WHERE delete_at IS NULL"
	rows, err := p_repo.DB.QueryContext(ctx, sql)
	product := modelProduct.Product{}
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.UserId, &product.Name, &product.Deskripsi)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return
}
func (p_repo *ProductRepoImpl) AdmUpdateProduct(ctx context.Context, productIn modelProduct.Product) (err error) {
	logCtx := fmt.Sprintf("%T - AdmupdateProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "UPDATE data_product SET name=$1,deskripsi=$2,update_at=$3 WHERE id=$4"
	_, err = p_repo.DB.ExecContext(ctx, sql, productIn.Name, productIn.Deskripsi, productIn.Update_at, productIn.Id)
	if err != nil {
		return
	}
	return
}
func (p_repo *ProductRepoImpl) AdmDeleteProduct(ctx context.Context, productId uint64, delete_at *time.Time) (err error) {
	logCtx := fmt.Sprintf("%T - AdmSoftdeleteProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "UPDATE data_product SET delete_at=$1 WHERE id=$2"
	_, err = p_repo.DB.ExecContext(ctx, sql, delete_at, productId)
	if err != nil {
		return
	}
	return
}

func (p_repo *ProductRepoImpl) UsrCreateProduct(ctx context.Context, productIn modelProduct.Product) (err error) {
	logCtx := fmt.Sprintf("%T - UsrCreateProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "INSERT INTO data_product(user_id,name,deskripsi,create_at) VALUES($1,$2,$3,$4)"
	_, err = p_repo.DB.ExecContext(ctx, sql, productIn.UserId, productIn.Name, productIn.Deskripsi, productIn.Create_at)
	if err != nil {
		return
	}
	return
}
func (p_repo *ProductRepoImpl) UsrFindAllProduct(ctx context.Context, userId int64) (products []modelProduct.Product, err error) {
	logCtx := fmt.Sprintf("%T - UsrFindAllProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,user_id,name,deskripsi FROM data_product WHERE user_id=$1 AND delete_at IS NULL"
	rows, err := p_repo.DB.QueryContext(ctx, sql, userId)
	product := modelProduct.Product{}
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.UserId, &product.Name, &product.Deskripsi)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return
}
func (p_repo *ProductRepoImpl) UsrFindByIdProduct(ctx context.Context, productId uint64, userId uint64) (product modelProduct.Product, err error) {
	logCtx := fmt.Sprintf("%T - usrFindbyIdProduct", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,user_id,name,deskripsi FROM data_product WHERE id=$1  AND user_id=$2 AND delete_at IS NULL"
	rows, err := p_repo.DB.QueryContext(ctx, sql, productId, userId)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&product.Id, &product.UserId, &product.Name, &product.Deskripsi)
		if err != nil {
			return
		}
		return
	} else {
		return product, errors.New("PRODUCT NOT FOUND")
	}
}
