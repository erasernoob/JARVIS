package service

import (
	"context"
	"fmt"
	"time"

	"github.com/erasernoob/JARVIS/common"
	db "github.com/erasernoob/JARVIS/dao"
	m "github.com/erasernoob/JARVIS/model"
	"github.com/google/uuid"
)

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
		// 直接返回所有conversations
		agent.Conversations = conversations
		agent.UID = uid
		agent.CurCID = conversations[len(conversations)-1].ID.String()
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
