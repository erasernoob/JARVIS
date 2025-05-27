package model

// 只做结构定义
import (
	"context"

	llm "github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type Client struct {
	UID           string
	LLM           llm.ToolCallingChatModel
	History       []*schema.Message
	Conversations []*Conversation
	CurCID        string // 现在正在进行着的conversation
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
	// Add Message

	return response.Content, nil
}

func (c *Client) SendSystemMessage(ctx context.Context, messages ...string) (string, error) {
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
