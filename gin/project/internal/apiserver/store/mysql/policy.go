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

type policy struct {
	db *gorm.DB
}

func newPolicy(ds *datastore) *policy {
	return &policy{db : ds.db}
}

func (p *policy) Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error {
	return p.db.Create(policy).Error
}


func (p *policy) Delete(ctx context.Context, username, name string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	err := p.db.Where("username = ? and name = ?", username, name).Delete(&v1.Policy{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Get return an user by the user identifier.
func (p *policy) Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error) {
	policy := &v1.Policy{}
	err := p.db.Where("username = ? and name = ?", username, name).First(&policy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrPolicyNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return policy, nil
}

func (p *policy) Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error {
	panic("implement me")
}

func (p *policy) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	panic("implement me")
}

func (p *policy) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error) {
	panic("implement me")
}
