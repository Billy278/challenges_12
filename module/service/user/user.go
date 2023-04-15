package user

import (
	"context"

	"github.com/Billy278/challenges_12-13/module/models/token"
	modelUser "github.com/Billy278/challenges_12-13/module/models/user"
)

type UserSrv interface {
	CreateUser(ctx context.Context, userIn modelUser.User) (err error)
	FindbyUsername(ctx context.Context, usernameIn string) (userRes modelUser.UserRes, err error)
	LoginByUserName(ctx context.Context, userIn modelUser.User) (tokens token.Tokens, err error)
}
