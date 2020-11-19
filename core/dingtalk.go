package core

import (
	"context"
	"net/http"
)

type IDingTalkClient interface {
	GetToken(ctx context.Context, req GetTokenRequest) (GetTokenResponse, error)
	ListUserId(ctx context.Context, req ListUserIdRequest) (ListUserIdResponse, error)
	GetUserInfo(ctx context.Context, req GetUserRequest) (GetUserResponse, error)
	SendCorpConversation(ctx context.Context, req SendCropConversationRequest) (SendCropConversationResponse, error)
	SendRobotMessage(ctx context.Context, sendRoboReq SendRobotMessageRequest) (resp SendRobotMessageResponse, err error)
}

type DingTalkClient struct {
	HttpClient *http.Client
}

func NewDingTalk(httpClient *http.Client) IDingTalkClient {
	return &DingTalkClient{HttpClient: httpClient}
}
