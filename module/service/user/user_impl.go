package user

import (
	"context"
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	token "github.com/Billy278/challenges_12-13/module/models/token"
	modelUser "github.com/Billy278/challenges_12-13/module/models/user"
	repository "github.com/Billy278/challenges_12-13/module/repository/user"
	helper "github.com/Billy278/challenges_12-13/pkg"
	cripto "github.com/Billy278/challenges_12-13/pkg/cripto"
)

type UserSrvImpl struct {
	UserRepo repository.UserRepo
}

func NewUserSrvImpl(userrepo repository.UserRepo) UserSrv {
	return &UserSrvImpl{
		UserRepo: userrepo,
	}
}

func (u_srv *UserSrvImpl) CreateUser(ctx context.Context, userIn modelUser.User) (err error) {
	_, err = u_srv.UserRepo.FindbyUsername(ctx, userIn.Username)
	if err == nil {
		log.Printf("[ERROR] error  Username Sudah ada :%v\n", err)
		return errors.New("USERNAME SUDAH ADA")
	}
	hashPass, err := cripto.GenerateHash(userIn.Password)
	if err != nil {
		log.Printf("[ERROR] error HashPassword :%v\n", err)
		return errors.New("HashPassword")
	}
	userIn.Password = hashPass
	tNow := time.Now()
	userIn.Create_at = &tNow
	err = u_srv.UserRepo.CreateUser(ctx, userIn)
	//book, err = booksrv_imp.RepoBook.CreateBook(ctx, bookIn)
	if err != nil {
		log.Printf("[ERROR] error Insert User :%v\n", err)
		return
	}
	return

}
func (u_srv *UserSrvImpl) FindbyUsername(ctx context.Context, usernameIn string) (userRes modelUser.UserRes, err error) {
	user, err := u_srv.UserRepo.FindbyUsername(ctx, usernameIn)
	if err != nil {
		log.Printf("[ERROR] error Find Username :%v\n", err)
		return
	}
	return helper.ToUserResponse(user), err
}

func (u_srv *UserSrvImpl) LoginByUserName(ctx context.Context, userIn modelUser.User) (tokens token.Tokens, err error) {
	user, err := u_srv.UserRepo.FindbyUsername(ctx, userIn.Username)
	if err != nil {
		log.Printf("[ERROR] error Find Username  :%v\n", err)
		return
	}
	// compare password
	// password acc -> hashed password
	// password login acc -> plain password
	err = cripto.CompareHash(user.Password, userIn.Password)
	if err != nil {
		log.Printf("[ERROR] PASSWORD IS VAILED  :%v\n", err)
		return
	}
	//create token
	jti := "jti"
	usr_id := strconv.Itoa(int(user.Id))
	idToken, accessToken, refreshToken, err := u_srv.generateAllTokensConcurency(ctx, usr_id, user.Username, user.Role, jti)
	if err != nil {
		return
	}

	return token.Tokens{
		IDToken:      idToken,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}

func (u_srv *UserSrvImpl) generateAllTokensConcurency(ctx context.Context, userid, username, role, jti string) (idToken, accessToken, refreshToken string, err error) {
	log.Printf("[INFO] Generate token  :%v\n", err)

	timeNow := time.Now()
	defaultClaim := token.DefaultClaim{
		Expired:   int(timeNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(timeNow.Unix()),
		IssuedAt:  int(timeNow.Unix()),
		Issuer:    userid,
		Audience:  "Challenges_12-13",
		JTI:       jti,
		Type:      token.ID_TOKEN,
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		//generate id token
		idTokenClaim := struct {
			token.DefaultClaim
			token.IDClaim
		}{
			defaultClaim_,
			token.IDClaim{
				Username: username,
				Role:     role,
			},
		}
		idToken, err = cripto.SignJWT(idTokenClaim)
		if err != nil {
			log.Printf("[ERROR] creating id token  :%v\n", err)
			return
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		//generate access token
		defaultClaim_.Expired = int(timeNow.Add(2 * time.Hour).Unix())
		defaultClaim_.Type = token.ACCESS_TOKEN
		accessTokenClaim := struct {
			token.DefaultClaim
			token.AccessClaim
		}{
			defaultClaim_,
			token.AccessClaim{
				Role:   role,
				UserID: userid,
			},
		}
		accessToken, err = cripto.SignJWT(accessTokenClaim)
		if err != nil {
			log.Printf("[ERROR] creating access token  :%v\n", err)
			return
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		//generate refresh token
		defaultClaim_.Expired = int(timeNow.Add(time.Hour).Unix())
		defaultClaim_.Type = token.REFRESH_TOKEN

		refreshTokenClaim := struct {
			token.DefaultClaim
		}{
			defaultClaim_,
		}
		refreshToken, err = cripto.SignJWT(refreshTokenClaim)
		if err != nil {
			log.Printf("[ERROR] creating refresh token  :%v\n", err)
			return
		}
	}(defaultClaim)

	wg.Wait()
	return
}
func (u_srv *UserSrvImpl) generateAllTokens(ctx context.Context, userid, username, role, jti string) (idToken, accessToken, refreshToken string, err error) {
	log.Printf("[INFO] Generate token  :%v\n", err)

	timeNow := time.Now()
	defaultClaim_ := token.DefaultClaim{
		Expired:   int(timeNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(timeNow.Unix()),
		IssuedAt:  int(timeNow.Unix()),
		Issuer:    userid,
		Audience:  "Challenges_12-13",
		JTI:       jti,
		Type:      token.ID_TOKEN,
	}

	//generate id token
	idTokenClaim := struct {
		token.DefaultClaim
		token.IDClaim
	}{
		defaultClaim_,
		token.IDClaim{
			Username: username,
			Role:     role,
		},
	}
	idToken, err = cripto.SignJWT(idTokenClaim)
	if err != nil {
		log.Printf("[ERROR] creating id token  :%v\n", err)
		return
	}

	//generate access token
	defaultClaim_.Expired = int(timeNow.Add(5 * time.Minute).Unix())
	defaultClaim_.Type = token.ACCESS_TOKEN
	accessTokenClaim := struct {
		token.DefaultClaim
		token.AccessClaim
	}{
		defaultClaim_,
		token.AccessClaim{
			Role:   role,
			UserID: userid,
		},
	}
	accessToken, err = cripto.SignJWT(accessTokenClaim)
	if err != nil {
		log.Printf("[ERROR] creating access token  :%v\n", err)
		return
	}

	//generate refresh token
	defaultClaim_.Expired = int(timeNow.Add(time.Hour).Unix())
	defaultClaim_.Type = token.REFRESH_TOKEN

	refreshTokenClaim := struct {
		token.DefaultClaim
	}{
		defaultClaim_,
	}
	refreshToken, err = cripto.SignJWT(refreshTokenClaim)
	if err != nil {
		log.Printf("[ERROR] creating refresh token  :%v\n", err)
		return
	}
	return
}
