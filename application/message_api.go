package application

import (
	"context"
	"github.com/lyouthzzz/dingtalk-go/core"
	"github.com/lyouthzzz/dingtalk-go/message"
	"github.com/pkg/errors"
)

// 消息api

func (c *Client) SendMarkDownMessage(ctx context.Context, title, text string,
	userIds []string) error {
	markDownMsg := message.MarkdownMessage{
		MsgType: message.MsgTypeMarkdown,
		Markdown: message.MarkdownParams{
			Title: title,
			Text:  text,
		},
	}
	return c.sendMessage(ctx, userIds, markDownMsg)
}

func (c *Client) SendTextMessage(ctx context.Context, title, text string,
	userIds []string) error {
	textMsg := message.TextMessage{
		MsgType: message.MsgTypeText,
		Text:    message.TextParams{Content: text},
	}
	return c.sendMessage(ctx, userIds, textMsg)
}

func (c *Client) sendMessage(ctx context.Context, userIds []string, msg interface{}) error {
	var (
		token       Token
		sendMsgReq  core.SendCropConversationRequest
		sendMsgResp core.SendCropConversationResponse
		err         error
	)
	if token, err = c.GetToken(ctx); err != nil {
		return err
	}
	sendMsgReq = core.SendCropConversationRequest{
		AgentId:     c.agentId,
		UserIdList:  userIds,
		ToAllUser:   false,
		Message:     msg,
		BaseRequest: core.BaseRequest{AccessToken: token.value},
	}
	if sendMsgResp, err = c.dingTalkClient.SendCorpConversation(ctx, sendMsgReq); err != nil {
		return err
	}
	if !sendMsgResp.Success() {
		return errors.New(sendMsgResp.ErrorMsg)
	}
	return nil
}
