package middleware

import (
	"net/http"
	"strings"

	tokenModel "github.com/Billy278/challenges_12-13/module/models/token"
	response "github.com/Billy278/challenges_12-13/pkg"
	"github.com/Billy278/challenges_12-13/pkg/cripto"
	"github.com/gin-gonic/gin"
)

type (
	HeaderKey  string
	ContextKey string
)

const (
	Authorization HeaderKey = "Authorization"

	AccessClaim ContextKey = "access_claim"

	BasicAuth  string = "Basic "
	BearerAuth string = "Bearer "
)

func BearerOAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// auth header
		header := ctx.GetHeader(string(Authorization))
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Token Not Found",
			})
			return
		}
		//get token
		token := strings.Split(header, BearerAuth)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Token Not Found",
			})
			return
		}

		// header token is found

		var claim tokenModel.AccessClaim
		err := cripto.ParseJWT(token[1], &claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "invalid token",
			})
			return
		}
		ctx.Set(string(AccessClaim), claim)
		ctx.Next()

	}

}
