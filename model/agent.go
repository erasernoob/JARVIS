package model

// 只做结构定义
import (
	"context"

	llm "github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type Client struct {
	LLM     llm.ToolCallingChatModel
	History []*schema.Message
	// 对话Id
	CID string
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

// func RestoreClientFromDB(ctx context.Context) (*Client, error) {
// 	var agent Client
// 	// 查看是否有对话历史
// 	uid := ctx.Value(common.UID).(string)

// 	conversations, err := db.GetConversationByUid(ctx, uid)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get conversations: %w", err)
// 	}
// 	if len(conversations) != 0 {

// 	}

// 	// 如果有对话历史，保存后返回

// 	// 如果没有对话历史，创建一个新的对话
// }
