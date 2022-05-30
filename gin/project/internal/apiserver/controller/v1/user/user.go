package user

import (
	"github.com/gin-gonic/gin"
	serv "github.com/joe/iam/internal/apiserver/service/v1"
	"github.com/joe/iam/internal/apiserver/store"
	"github.com/joe/iam/internal/pkg/code"
	v1 "github.com/marmotedu/api/apiserver/v1"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
)

type UserController struct {
	srv serv.Service
}

func (c UserController) Get(context *gin.Context) {
	user, err := c.srv.Users().Get(context, context.Param("name"), metav1.GetOptions{})

	if err != nil {
		core.WriteResponse(context, err, nil)
		return
	}

	core.WriteResponse(context, nil, user)
}

func (u UserController) Create(c *gin.Context) {
	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	if err := u.srv.Users().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, r)

}

func (u *UserController) Delete (c *gin.Context) {
	err := u.srv.Users().Delete(c, c.Param("name"), metav1.DeleteOptions{Unscoped: true})

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, "success")
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{srv: serv.NewService(store)}
}
