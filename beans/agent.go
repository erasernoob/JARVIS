package beans

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type Client struct {
	LLM     model.ToolCallingChatModel
	History []*schema.Message
	// 对话Id
	CID string
}

// test
func (c *Client) SendUserMessage(ctx context.Context, messages ...string) (string, error) {
	// 首先先查看改user是否有历史对话记录，直接返回最后一个conversation
	// GetConversation(ctx)

	for _, msg := range messages {
		c.History = append(c.History, schema.UserMessage(msg))
	}

	response, err := c.LLM.Generate(ctx, c.History)
	if err != nil {
		return "", err
	}
	// Add Message

	return response.Content, nil
}

func (c *Client) SenSystemMessage(ctx context.Context, messages ...string) (string, error) {
	for _, msg := range messages {
		c.History = append(c.History, schema.SystemMessage(msg))
	}

	response, err := c.LLM.Generate(ctx, c.History)
	if err != nil {
		return "", err
	}
	// Add Message

	return response.Content, nil
}
