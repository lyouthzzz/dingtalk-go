package core

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/lyouthzzz/dingtalk-go/util"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 工作通知消息
func (c *DingTalkClient) SendCorpConversation(ctx context.Context, req SendCropConversationRequest) (resp SendCropConversationResponse, err error) {
	param := url.Values{}
	param.Add(ConstAccessToken, req.AccessToken)
	sendUrl := SendCropConversationUrl + "?" + param.Encode()

	bodyMap := make(map[string]interface{})
	bodyMap[ConstAgentId] = req.AgentId
	bodyMap[ConstUserIdList] = strings.Join(req.UserIdList, ",")
	if req.DeptIdList != nil && len(req.DeptIdList) != 0 {
		bodyMap[ConstToAllUser] = strings.Join(req.DeptIdList, ",")
	}
	bodyMap[ConstToAllUser] = req.ToAllUser
	bodyMap[ConstMsg] = req.Message
	bodyBytes, err := jsoniter.Marshal(bodyMap)
	if err != nil {
		return
	}
	httpOption := util.HttpOption{
		Client: c.HttpClient,
		Ctx:    ctx,
		Body:   bytes.NewReader(bodyBytes),
		Header: http.Header{
			"Content-type": {"application/json"},
		},
	}
	if err = util.DoHttpRequest(sendUrl, http.MethodPost, &resp, httpOption); err != nil {
		return
	}
	return
}

func (c *DingTalkClient) SendRobotMessage(ctx context.Context, sendRoboReq SendRobotMessageRequest) (resp SendRobotMessageResponse, err error) {
	var (
		bodyBytes       []byte
		param           url.Values
		sendRobotReqUrl string
		genAuthParam    = func(secret string) (url.Values, error) {
			timestamp := time.Now().Unix() * 1000
			strToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
			h := hmac.New(sha256.New, []byte(secret))
			if _, err := h.Write([]byte(strToSign)); err != nil {
				return nil, err
			}
			sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
			param := url.Values{}
			param.Add("timestamp", fmt.Sprintf("%d", timestamp))
			param.Add("sign", sign)
			return param, nil
		}
	)
	if param, err = genAuthParam(sendRoboReq.Secret); err != nil {
		return
	}
	param.Add(ConstAccessToken, sendRoboReq.Token)
	// https://oapi.dingtalk.com/robot/send?access_token=
	sendRobotReqUrl = SendRobotMessageUrl + "?" + param.Encode()
	if bodyBytes, err = jsoniter.Marshal(sendRoboReq.Message); err != nil {
		return
	}
	httpOption := util.HttpOption{
		Client: c.HttpClient,
		Ctx:    ctx,
		Body:   bytes.NewReader(bodyBytes),
		Header: http.Header{"Content-Type": {"application/json"}},
	}
	if err = util.DoHttpRequest(sendRobotReqUrl, http.MethodPost, &resp, httpOption); err != nil {
		return
	}
	if !resp.Success() {
		err = errors.New(resp.ErrorMsg)
		return
	}
	return
}
