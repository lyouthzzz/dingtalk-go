package application

import (
	"context"
	"github.com/lyouthzzz/dingtalk-go/core"
	"github.com/pkg/errors"
)

// 通讯录API

func (c *Client) ListUserId(ctx context.Context, deptId uint32) ([]string, error) {
	var (
		token          Token
		listUserIdReq  core.ListUserIdRequest
		listUserIdResp core.ListUserIdResponse
		err            error
	)
	if token, err = c.GetToken(ctx); err != nil {
		return nil, err
	}
	listUserIdReq = core.ListUserIdRequest{
		BaseRequest: core.BaseRequest{AccessToken: token.value},
		DeptId:      deptId,
	}
	if listUserIdResp, err = c.dingTalkClient.ListUserId(ctx, listUserIdReq); err != nil {
		return nil, err
	}
	if !listUserIdResp.Success() {
		return nil, errors.New(listUserIdResp.ErrorMsg)
	}
	return listUserIdResp.UserIds, nil
}

func (c *Client) GetUserInfo(ctx context.Context, useId string) (userInfo UserInfo, err error) {
	var (
		token           Token
		getUserInfoReq  core.GetUserRequest
		getUserInfoResp core.GetUserResponse
	)
	if token, err = c.GetToken(ctx); err != nil {
		return
	}
	getUserInfoReq = core.GetUserRequest{
		BaseRequest: core.BaseRequest{AccessToken: token.value},
		UserId:      useId,
	}
	if getUserInfoResp, err = c.dingTalkClient.GetUserInfo(ctx, getUserInfoReq); err != nil {
		return
	}
	if !getUserInfoResp.Success() {
		err = errors.New(getUserInfoResp.ErrorMsg)
		return
	}
	userInfo = userInfo.From(getUserInfoResp)
	return
}
