package product

import (
	"context"
	"errors"
	"testing"
	"time"

	modelProduct "github.com/Billy278/challenges_12-13/module/models/user"
	repomock "github.com/Billy278/challenges_12-13/module/repository/product/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProductFound(t *testing.T) {
	createdAt := time.Now()
	type (
		input struct {
			productIn modelProduct.Product
		}
		want struct {
			productRes []modelProduct.ProductRes
			err        error
		}
	)
	testCases := []struct {
		desc string
		// dia akan mocking seakan akan function yang kita test menjalankan procedure function kita
		doMock func(repoMock *repomock.MockProductRepo)
		input  input
		want   want
	}{
		{
			desc: "happy case",
			input: input{
				productIn: modelProduct.Product{
					Id:        1,
					UserId:    1,
					Name:      "pensil",
					Deskripsi: "alat tulis",
					Create_at: &createdAt,
				},
			}, want: want{
				err: nil,
				productRes: []modelProduct.ProductRes{
					{
						Id:        1,
						UserId:    1,
						Name:      "pensil",
						Deskripsi: "alat tulis",
					},
				},
			},
			doMock: func(repoMock *repomock.MockProductRepo) {
				repoMock.EXPECT().AdmFindAllProduct(gomock.Any()).Return(
					[]modelProduct.Product{{
						Id:        1,
						UserId:    1,
						Name:      "pensil",
						Deskripsi: "alat tulis",
					},
					}, nil).
					MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repomock := repomock.NewMockProductRepo(ctrl)
			tC.doMock(repomock)
			svc := ProductSrvImpl{
				ProductRepo: repomock,
			}
			findAll, _ := svc.AdmFindAllProductSrv(context.Background())
			assert.Equal(t, tC.want.productRes, findAll)
		})
	}
}

// tes get all product not found
func TestGetAllProductNotFound(t *testing.T) {

	type (
		want struct {
			productRes []modelProduct.ProductRes
			err        error
		}
	)
	testCases := []struct {
		desc string
		// dia akan mocking seakan akan function yang kita test menjalankan procedure function kita
		doMock func(repoMock *repomock.MockProductRepo)
		want   want
	}{
		{
			desc: "happy case",
			want: want{
				err:        errors.New("NOT FOUND"),
				productRes: []modelProduct.ProductRes{},
			},
			doMock: func(repoMock *repomock.MockProductRepo) {
				repoMock.EXPECT().AdmFindAllProduct(gomock.Any()).Return(
					[]modelProduct.Product{}, errors.New("NOT FOUND")).
					MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repomock := repomock.NewMockProductRepo(ctrl)
			tC.doMock(repomock)
			svc := ProductSrvImpl{
				ProductRepo: repomock,
			}
			_, err := svc.AdmFindAllProductSrv(context.Background())

			assert.Equal(t, tC.want.err, err)
		})
	}
}

// tes get one product by id  found
func TestGetOneProductFound(t *testing.T) {
	createdAt := time.Now()
	type (
		input struct {
			productIn modelProduct.Product
		}
		want struct {
			productRes modelProduct.ProductRes
			err        error
		}
	)
	testCases := []struct {
		desc string
		// dia akan mocking seakan akan function yang kita test menjalankan procedure function kita
		doMock func(repoMock *repomock.MockProductRepo)
		input  input
		want   want
	}{
		{
			desc: "happy case",
			input: input{
				productIn: modelProduct.Product{
					Id:        1,
					UserId:    1,
					Name:      "pensil",
					Deskripsi: "alat tulis",
					Create_at: &createdAt,
				},
			}, want: want{
				err: nil,
				productRes: modelProduct.ProductRes{
					Id:        1,
					UserId:    1,
					Name:      "pensil",
					Deskripsi: "alat tulis",
				},
			},
			doMock: func(repoMock *repomock.MockProductRepo) {
				repoMock.EXPECT().AdmFindbyIdProduct(gomock.Any(), uint64(1)).Return(
					modelProduct.Product{
						Id:        1,
						UserId:    1,
						Name:      "pensil",
						Deskripsi: "alat tulis",
					}, nil).
					MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repomock := repomock.NewMockProductRepo(ctrl)
			tC.doMock(repomock)
			svc := ProductSrvImpl{
				ProductRepo: repomock,
			}
			findAll, _ := svc.AdmFindbyIdProuctSrv(context.Background(), uint64(1))
			assert.Equal(t, tC.want.productRes, findAll)
		})
	}
}

// tes get one product by id Not  found
func TestGetOneProductNotFound(t *testing.T) {

	type (
		want struct {
			productRes modelProduct.ProductRes
			err        error
		}
	)
	testCases := []struct {
		desc string
		// dia akan mocking seakan akan function yang kita test menjalankan procedure function kita
		doMock func(repoMock *repomock.MockProductRepo)

		want want
	}{
		{
			desc: "happy case",
			want: want{
				err:        errors.New("NOT FOUND"),
				productRes: modelProduct.ProductRes{},
			},
			doMock: func(repoMock *repomock.MockProductRepo) {
				repoMock.EXPECT().AdmFindbyIdProduct(gomock.Any(), uint64(1)).Return(
					modelProduct.Product{}, errors.New("NOT FOUND")).
					MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repomock := repomock.NewMockProductRepo(ctrl)
			tC.doMock(repomock)
			svc := ProductSrvImpl{
				ProductRepo: repomock,
			}
			findAll, err := svc.AdmFindbyIdProuctSrv(context.Background(), uint64(1))
			assert.Equal(t, tC.want.productRes, findAll)
			assert.Equal(t, tC.want.err, err)
		})
	}
}

//mockgen -source=module/repository/product/product.go -destination=module/repository/product/mock/product_mock.go -package=mock
