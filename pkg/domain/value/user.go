package value

import "errors"

type (
	UserName  string
	UserScore int
)

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
		return errors.New("user name is empty")
	}
	return nil
}

func NewScore(score int) UserScore {
	return UserScore(score)
}

func (s UserScore) Int() int {
	return int(s)
}

func (s UserScore) Set(score int) UserScore {
	s = UserScore(score)
	return s
}

func (s UserScore) Validate() error {
	if s < 0 {
		return errors.New("score is negative")
	}
	return nil
}
