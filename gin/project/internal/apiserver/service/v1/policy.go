package v1

import (
	"context"
	store2 "github.com/joe/iam/internal/apiserver/store"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type PolicySrv interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
	Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}


type policyService struct {
	store store2.Factory
}

func (p policyService) Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error {
	return p.store.Policies().Create(ctx, policy, opts)
}

func (p policyService) Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error {
	panic("implement me")
}

func (p policyService) Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error {
	return p.store.Policies().Delete(ctx, username, name, opts)
}

func (p policyService) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	panic("implement me")
}

func (p policyService) Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error) {
	return p.store.Policies().Get(ctx, username, name, opts)
}

func (p policyService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error) {
	panic("implement me")
}

var _ PolicySrv = (*policyService)(nil)

func NewPolicy(store store2.Factory) PolicySrv{
	return &policyService{
		store: store,
	}
}