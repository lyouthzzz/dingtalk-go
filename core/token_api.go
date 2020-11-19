package core

import (
	"context"
	"github.com/lyouthzzz/dingtalk-go/util"
	"net/http"
	"net/url"
)

func (c *DingTalkClient) GetToken(ctx context.Context, getTokenReq GetTokenRequest) (resp GetTokenResponse, err error) {
	param := url.Values{}
	param.Add(ConstAppKey, getTokenReq.AppKey)
	param.Add(ConstAppSecret, getTokenReq.AppSecret)
	getTokenReqUrl := GetTokenRequestUrl + "?" + param.Encode()
	httpOption := util.HttpOption{
		Client: c.HttpClient,
		Ctx:    ctx,
	}
	if err = util.DoHttpRequest(getTokenReqUrl, http.MethodGet, &resp, httpOption); err != nil {
		return
	}
	return
}
