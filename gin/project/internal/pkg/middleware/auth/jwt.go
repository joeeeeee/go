package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joe/iam/internal/pkg/middleware"
)

const AuthzAudience = "iam.authz.joe.com"

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ middleware.AuthStrategy = &JWTStrategy{}

func NewJWTStrategy(jwtMiddleware ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{jwtMiddleware}
}

func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}
