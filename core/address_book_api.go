package core

import (
	"context"
	"github.com/lyouthzzz/dingtalk-go/util"
	"net/http"
	"net/url"
	"strconv"
)

func (c *DingTalkClient) ListUserId(ctx context.Context, req ListUserIdRequest) (resp ListUserIdResponse, err error) {
	param := url.Values{}
	param.Add(ConstAccessToken, req.AccessToken)
	param.Add(ConstDeptId, strconv.FormatInt(int64(req.DeptId), 10))
	listUserReqUrl := ListUserIdUrl + "?" + param.Encode()

	httpOption := util.HttpOption{
		Client: c.HttpClient,
		Ctx:    ctx,
	}
	if err = util.DoHttpRequest(listUserReqUrl, http.MethodGet, &resp, httpOption); err != nil {
		return
	}
	return
}

func (c *DingTalkClient) GetUserInfo(ctx context.Context, req GetUserRequest) (resp GetUserResponse, err error) {
	param := url.Values{}
	param.Add(ConstUserId, req.UserId)
	param.Add(ConstAccessToken, req.AccessToken)
	getUserUrl := GetUserInfoUrl + "?" + param.Encode()
	httpOption := util.HttpOption{
		Client: c.HttpClient,
		Ctx:    ctx,
	}
	if err = util.DoHttpRequest(getUserUrl, http.MethodGet, &resp, httpOption); err != nil {
		return
	}
	return
}
