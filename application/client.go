package application

import (
	"github.com/lyouthzzz/dingtalk-go/core"
	"net/http"
	"sync"
)

// 小程序应用

type Option struct {
	Key     string
	Secret  string
	AgentId uint32
}

type Client struct {
	key, secret    string
	agentId        uint32
	mu             sync.Mutex
	token          Token
	dingTalkClient core.IDingTalkClient
}

func NewClient(opts Option) *Client {
	return &Client{
		agentId:        opts.AgentId,
		key:            opts.Key,
		secret:         opts.Secret,
		dingTalkClient: core.NewDingTalk(http.DefaultClient),
	}
}
