package service

import (
	rest_rr "github.com/marcossdev04/crud-go/src/configuration/rest_err"
	"github.com/marcossdev04/crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_rr.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_rr.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_rr.RestErr)
	DeleteUser(string) *rest_rr.RestErr
}
