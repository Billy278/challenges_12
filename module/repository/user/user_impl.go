package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	modelUser "github.com/Billy278/challenges_12-13/module/models/user"
)

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}

func (u_repo *UserRepoImpl) CreateUser(ctx context.Context, userIn modelUser.User) (err error) {
	logCtx := fmt.Sprintf("%T - CreateAccount", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "INSERT INTO data_user(name,username,password,role,create_at) VALUES($1,$2,$3,$4,$5)"
	_, err = u_repo.DB.ExecContext(ctx, sql, userIn.Name, userIn.Username, userIn.Password, userIn.Role, userIn.Create_at)
	if err != nil {
		return
	}
	return

}

func (u_repo *UserRepoImpl) FindbyUsername(ctx context.Context, usernameIn string) (user modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - FindUser", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,name,username,password,role FROM data_user WHERE username=$1"
	rows, err := u_repo.DB.QueryContext(ctx, sql, usernameIn)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Role)

		return
	} else {
		return user, errors.New("USER IS NOT FOUND")
	}
}
