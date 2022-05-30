package apiserver

import (
	"encoding/base64"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joe/iam/internal/apiserver/store"
	"github.com/joe/iam/internal/pkg/middleware"
	"github.com/joe/iam/internal/pkg/middleware/auth"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/spf13/viper"
	"google.golang.org/appengine/log"
	"strings"
	"time"
)

var identityKey = "id"

type loginInfo struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func newJWTAuth() middleware.AuthStrategy {
	middleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       viper.GetString("jwt.Realm"),
		Key:         []byte(viper.GetString("jwt.key")),
		Timeout:     viper.GetDuration("jwt.timeout"),
		MaxRefresh:  viper.GetDuration("jwt.max-refresh"),
		IdentityKey: middleware.UsernameKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			claims := jwt.MapClaims{
				"iss": nil,
				"aud": nil,
			}
			if u, ok := data.(*v1.User); ok {
				claims[jwt.IdentityKey] = u.Name
				claims["sub"] = u.Name
			}
			return claims
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return claims[jwt.IdentityKey]
		},
		Authenticator: authenticator(),
		Authorizator:  authorizator(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	return auth.NewJWTStrategy(*middleware)
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if _, ok := data.(string); ok {
			return true
		}
		return false
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals loginInfo
		var err error

		// token 方式请求
		if c.Request.Header.Get("Authorization") != "" {
			loginVals, err = parseWithToken(c)
		} else {
			loginVals, err = parseWithBody(c)
		}

		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		user, err := store.Client().Users().Get(c, loginVals.UserName, metav1.GetOptions{})

		if err != nil {
			log.Errorf(c, "get user information failed:%s", err.Error())

			return "", jwt.ErrFailedAuthentication
		}

		if err := user.Compare(loginVals.Password); err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		return user, nil
	}
}

func parseWithBody(c *gin.Context) (loginInfo, error) {
	var login loginInfo
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Errorf(c, "auth parse error: %s", err.Error())
		return loginInfo{}, jwt.ErrFailedAuthentication
	}
	return login, nil
}

func parseWithToken(c *gin.Context) (loginInfo, error) {

	// 获取token
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		log.Errorf(c, "get %s string from Authorization header failed", "Basic")
		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	// base64 解析token 信息
	payload, err := base64.StdEncoding.DecodeString(auth[1])

	if err != nil {
		log.Errorf(c, "decode basic string %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		log.Errorf(c, "parse payload failed", "")
		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return loginInfo{
		UserName: pair[0],
		Password: pair[1],
	}, nil

}
