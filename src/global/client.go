package global

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type Client struct {
	LLM     model.ToolCallingChatModel
	History []*schema.Message
}

// test
func (c *Client) SendUserMessage(ctx context.Context, messages ...string) (string, error) {
	for _, msg := range messages {
		c.History = append(c.History, schema.UserMessage(msg))
	}
	response, err := c.LLM.Generate(ctx, c.History)
	if err != nil {
		return "", err
	}
	return response.Content, nil
}
