package v1

import (
	"context"
	store2 "github.com/joe/iam/internal/apiserver/store"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type SecretSrv interface {
	Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error
	Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username, secretID string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, secretIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*v1.Secret, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.SecretList, error)
}

type secretService struct {
	store store2.Factory
}

func (s secretService) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error {
	panic("implement me")
}

func (s secretService) Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) error {
	panic("implement me")
}

func (s secretService) Delete(ctx context.Context, username, secretID string, opts metav1.DeleteOptions) error {
	panic("implement me")
}

func (s secretService) DeleteCollection(ctx context.Context, username string, secretIDs []string, opts metav1.DeleteOptions) error {
	panic("implement me")
}

func (s secretService) Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*v1.Secret, error) {
	panic("implement me")
}

func (s secretService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.SecretList, error) {
	panic("implement me")
}

func NewSecret(store store2.Factory) SecretSrv{
	return &secretService{
		store: store,
	}
}