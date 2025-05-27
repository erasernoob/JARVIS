package service

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/common"
	"github.com/erasernoob/JARVIS/dao"
	db "github.com/erasernoob/JARVIS/dao"
	"github.com/erasernoob/JARVIS/model"
	m "github.com/erasernoob/JARVIS/model"
	"github.com/google/uuid"
)

func RestoreTheChatHistory(ctx context.Context, agent *m.Client) error {
	// 1. 根据当前活跃conversation Id 拿到messages
	var msgs []*schema.Message
	msgs, err := dao.GetMessagesByConversationID(ctx, agent.CurCID)
	if err != nil {
		return err
	}
	slices.Reverse(msgs)
	agent.History = append(agent.History, msgs...)
	return nil
}

func RestoreClientFromDB(ctx context.Context) (*m.Client, error) {

	var agent m.Client
	// 查看是否有对话历史
	uid := ctx.Value(common.UID).(string)

	conversations, err := db.GetConversationByUid(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversations: %w", err)
	}

	// 如果有对话历史，保存后返回
	if len(conversations) != 0 {
		agent.Conversations = conversations
		agent.UID = uid
		agent.CurCID = conversations[len(conversations)-1].ID.String()

		// 2.Get the messages
		if err := RestoreTheChatHistory(ctx, &agent); err != nil {
			return nil, fmt.Errorf("restore the chat history failed: %v", err)
		}

		return &agent, nil
	}
	// 如果没有对话历史，创建一个新的对话
	cid, _ := uuid.NewRandom()
	title := fmt.Sprintf("%s-%s", uid, cid.String())

	now := time.Now()

	newConversation := &m.Conversation{
		ID:        cid,
		Title:     title,
		UserID:    uid,
		CreatedAt: now,
		UpdatedAt: now,
	}

	dbId, err := db.AddConversation(ctx, newConversation)
	fmt.Println(dbId)
	if err != nil {
		return nil, fmt.Errorf("failed to create new conversation: %w", err)
	}
	agent.Conversations = append(agent.Conversations, newConversation)
	agent.UID = uid
	// 设置活跃的conversation
	agent.CurCID = newConversation.ID.String()
	return &agent, nil
}

// 对外暴露发送消息的接口
func SendUserMessage(ctx context.Context, agent *model.Client, messages ...string) (string, error) {
	response, err := agent.SendUserMessage(ctx, messages...)
	if err != nil {
		return "", fmt.Errorf("agent send the user message failed: %w", err)
	}

	// persist the query
	if err = dao.AddMessage(ctx, agent.CurCID, common.USER, messages[0]); err != nil {
		return "", fmt.Errorf("persist the message failed: %w", err)
	}

	// persist the response and the messages
	if err = dao.AddMessage(ctx, agent.CurCID, common.ASSISTANT, response.Content); err != nil {
		return "", fmt.Errorf("persist the message failed: %w", err)
	}

	return response.Content, nil
}
