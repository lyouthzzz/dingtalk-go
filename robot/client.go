package robot

import (
	"github.com/lyouthzzz/dingtalk-go/core"
	"net/http"
)

// 群机器人

type Option struct {
	AccessToken string
	Secret      string
}

type Client struct {
	accessToken    string
	secret         string
	dingTalkClient core.IDingTalkClient
}

func NewClient(opts Option) *Client {
	return &Client{
		accessToken:    opts.AccessToken,
		secret:         opts.Secret,
		dingTalkClient: core.NewDingTalk(http.DefaultClient),
	}
}
