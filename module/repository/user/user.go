package user

import (
	"context"

	modelUser "github.com/Billy278/challenges_12-13/module/models/user"
)

type UserRepo interface {
	CreateUser(ctx context.Context, userIn modelUser.User) (err error)
	FindbyUsername(ctx context.Context, usernameIn string) (user modelUser.User, err error)
}
