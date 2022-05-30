// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"github.com/joe/iam/internal/pkg/code"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"gorm.io/gorm"
)

// UserStore defines the user storage interface.

type users struct {
	db *gorm.DB
}

func (u *users) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	return u.db.Create(user).Error
}

func (u *users) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	return u.db.Create(&user).Error
}

func (u *users) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	err := u.db.Where("name = ?", username).Delete(&v1.User{}).Error

	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil

}

func (u *users) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	panic("implement me")
}

func (u *users) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	panic("implement me")
}

func newUsers(ds *datastore) *users {
	return &users{db : ds.db}
}

// Get return an user by the user identifier.
func (u *users) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return user, nil
}
