package v1

import (
	store2 "github.com/joe/iam/internal/apiserver/store"
)

type Service interface {
	Users() UserSrv
	Secrets() SecretSrv
	Policies() PolicySrv
}

type service struct {
	store store2.Factory
}

func NewService(store store2.Factory) Service {
	return &service{store}
}

func (s service) Users() UserSrv {
	return NewUser(s.store)
}

func (s service) Secrets() SecretSrv {
	panic("implement me")
}

func (s service) Policies() PolicySrv {
	return NewPolicy(s.store)
}





