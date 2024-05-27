package model

import (
	"crypto/md5"
	"encoding/hex"

	rest_rr "github.com/marcossdev04/crud-go/src/configuration/rest_err"
)

type UserDomainInterface interface {
	GetAge() int8
	GetPassword() string
	GetName() string
	GetEmail() string

	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{
		email, name, password, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}
func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomain interface {
	CreateUser() *rest_rr.RestErr
	UpdateUser(string) *rest_rr.RestErr
	FindUser(string) (*userDomain, rest_rr.RestErr)
	DeleteUser(string) *rest_rr.RestErr
}
