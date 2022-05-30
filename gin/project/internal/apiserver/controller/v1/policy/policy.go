package policy

import (
	"github.com/gin-gonic/gin"
	serv "github.com/joe/iam/internal/apiserver/service/v1"
	"github.com/joe/iam/internal/apiserver/store"
	"github.com/joe/iam/internal/pkg/code"
	"github.com/joe/iam/internal/pkg/middleware"
	v1 "github.com/marmotedu/api/apiserver/v1"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
)

type PolicyController struct {
	srv serv.Service
}

func NewPolicyController(store store.Factory) *PolicyController {
	return &PolicyController{srv: serv.NewService(store)}
}

func (p *PolicyController) Create(c *gin.Context) {

	var r v1.Policy
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	r.Username = c.GetString(middleware.UsernameKey)

	if err := p.srv.Policies().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, r)
}

func (p *PolicyController) Get(c *gin.Context) {
	pol, err := p.srv.Policies().Get(c, c.GetString(middleware.UsernameKey), c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, pol)
}

func (p *PolicyController) Delete(c *gin.Context) {
	if err := p.srv.Policies().Delete(c, c.GetString(middleware.UsernameKey), c.Param("name"),
		metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}