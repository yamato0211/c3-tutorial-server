package usermodel

import "github.com/pkg/errors"

type (
	UserName  string
	UserScore int
)

var userNameError = errors.New("user name is empty")

func NewUserName(name string) UserName {
	return UserName(name)
}

func (n UserName) String() string {
	return string(n)
}

func (n UserName) Set(name string) UserName {
	n = UserName(name)
	return n
}

func (u UserName) Validate() error {
	if len(u) == 0 {
		return errors.WithStack(userNameError)
	}
	return nil
}
