package robot

import (
	"context"
	"github.com/lyouthzzz/dingtalk-go/core"
	"github.com/lyouthzzz/dingtalk-go/message"
	"github.com/pkg/errors"
)

func (c *Client) SendMarkDownMessage(ctx context.Context, title, text string,
	atMobiles []string, isAtAll bool) error {
	var (
		markdownMessage message.MarkdownMessage
		sendRobotReq    core.SendRobotMessageRequest
		sendRobotResp   core.SendRobotMessageResponse
		err             error
	)
	markdownMessage = message.MarkdownMessage{
		MsgType: message.MsgTypeMarkdown,
		Markdown: message.MarkdownParams{
			Title: title,
			Text:  text,
		},
		At: message.AtParams{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	}
	sendRobotReq = core.SendRobotMessageRequest{
		Token:   c.accessToken,
		Secret:  c.secret,
		Message: markdownMessage,
	}
	if sendRobotResp, err = c.dingTalkClient.SendRobotMessage(ctx, sendRobotReq); err != nil {
		return err
	}
	if !sendRobotResp.Success() {
		return errors.New(sendRobotResp.ErrorMsg)
	}
	return nil
}
