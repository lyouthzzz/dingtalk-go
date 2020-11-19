package application

import (
	"context"
	"github.com/lyouthzzz/dingtalk-go/core"
	"github.com/pkg/errors"
	"time"
)

func (c *Client) GetToken(ctx context.Context) (token Token, err error) {
	c.mu.Lock()
	c.mu.Unlock()
	if c.token.IsExpire() {
		if _, err = c.refreshToken(ctx); err != nil {
			return
		}
	}
	return c.token, nil
}

func (c *Client) refreshToken(ctx context.Context) (token Token, err error) {
	var (
		getTokenReq = core.GetTokenRequest{
			AppKey:    c.key,
			AppSecret: c.secret,
		}
		getTokenResp core.GetTokenResponse
	)

	if getTokenResp, err = c.dingTalkClient.GetToken(ctx, getTokenReq); err != nil {
		return
	}
	if !getTokenResp.Success() {
		err = errors.New(getTokenResp.ErrorMsg)
		return
	}
	token.expireAt = time.Now().Add(time.Second * time.Duration(getTokenResp.ExpiresIn))
	token.value = getTokenResp.AccessToken
	c.token = token
	return
}
