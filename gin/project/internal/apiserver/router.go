package apiserver

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joe/iam/internal/apiserver/controller/v1/policy"
	"github.com/joe/iam/internal/apiserver/controller/v1/user"
	"github.com/joe/iam/internal/apiserver/store/mysql"
	auth "github.com/joe/iam/internal/pkg/middleware/auth"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	gindump "github.com/tpkeeper/gin-dump"
	"log"
	"net/http"
)

func initRouter() {
	installController()
}

func installController() {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	jwtStrategy := newJWTAuth().(auth.JWTStrategy)

	errorInit := jwtStrategy.MiddlewareInit()

	if errorInit != nil {
		log.Fatal("auth.middlewareInit() error :" + errorInit.Error())
	}

	r.POST("/login", jwtStrategy.LoginHandler)
	r.POST("/logout", jwtStrategy.LogoutHandler)

	r.NoRoute(jwtStrategy.AuthFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	storeIns, _ := mysql.GetMySQLFactoryOr(nil)

	v1 := r.Group("/v1", gindump.Dump())
	{
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)
			userv1.GET(":name", userController.Get)
			userv1.POST("", userController.Create)
			userv1.DELETE(":name", userController.Delete)
		}

		policyv1 := v1.Group("/policy",jwtStrategy.AuthFunc())
		{
			policyController := policy.NewPolicyController(storeIns)
			policyv1.GET(":name", policyController.Get)
			policyv1.POST("", policyController.Create)
			policyv1.DELETE(":name", policyController.Delete)
		}
	}

	a := r.Group("/auth")

	a.GET("/refresh_token", jwtStrategy.RefreshHandler)

	a.Use(jwtStrategy.MiddlewareFunc())
	{
		a.GET("/hello", helloHandler)
	}

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}

func helloHandler (c *gin.Context) {
	storeIns, _ := mysql.GetMySQLFactoryOr(nil)

	user, err := storeIns.Users().Get(c, "admin", metav1.GetOptions{})

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	//claims := jwt.ExtractClaims(c)
	//user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userName" : user.Name,
		"instantId" : user.InstanceID,
		"text" : "hello world",
	})
}
