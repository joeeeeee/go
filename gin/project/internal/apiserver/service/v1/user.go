package v1

import (
	"context"
	store2 "github.com/joe/iam/internal/apiserver/store"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ChangePassword(ctx context.Context, user *v1.User) error
}

type userService struct {
	store store2.Factory
}

func NewUser(store store2.Factory) userService {
	return userService{
		store: store,
	}
}

var _ UserSrv = (*userService)(nil)

func (u userService) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	return u.store.Users().Create(ctx, user,opts)
}

func (u userService) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	panic("implement me")
}

func (u userService) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	return u.store.Users().Delete(ctx, username, opts)
}

func (u userService) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	panic("implement me")
}

func (u userService) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	return u.store.Users().Get(ctx,username,opts)
}

func (u userService) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	panic("implement me")
}

func (u userService) ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	panic("implement me")
}

func (u userService) ChangePassword(ctx context.Context, user *v1.User) error {
	panic("implement me")
}


