package application

import (
	"github.com/lyouthzzz/dingtalk-go/core"
	"time"
)

type Token struct {
	value    string
	expireAt time.Time
}

func (token Token) IsExpire() bool {
	return token.expireAt.IsZero() || token.expireAt.After(time.Now())
}

type UserInfo struct {
	UnionId string `json:"unionid"`
	UserId  string `json:"userid"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

func (u UserInfo) From(resp core.GetUserResponse) UserInfo {
	u.UnionId = resp.UnionId
	u.UserId = resp.UserId
	u.Mobile = resp.Mobile
	u.Email = resp.Email
	u.Name = resp.Name
	return u
}
