package mem

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/dao"
	"github.com/erasernoob/JARVIS/model"
	"github.com/google/uuid"
)

type Memory interface {
	// GetCurrentConversation
	SetCurConversation(conversation *model.Conversation) error
	GetCurConversation() (*model.Conversation, error)
	GetHistory() ([]*schema.Message, error)

	StoreMessageToDao(ctx context.Context, role schema.RoleType, message string) error
	AppendMessage(ctx context.Context, role schema.RoleType, message string) error
}

func NewMemoryMgr(uid string) (*MemoryMgr, error) {
	mgr := &MemoryMgr{
		curCvs:  nil,
		history: make([]*schema.Message, 0),
	}
	mgr.mux = sync.Mutex{}
	conversations, err := dao.GetConversationByUid(context.Background(), uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversation by uid: %w", err)
	}
	// choose the latest conversation
	if len(conversations) > 0 {
		mgr.curCvs = conversations[len(conversations)-1]
		// load the history
		history, err := dao.GetMessagesByConversationID(context.Background(), mgr.curCvs.ID.String())
		if err != nil {
			return nil, err
		}
		mgr.history = append(mgr.history, history...)
		fmt.Printf("Loaded history for conversation: %v\n", mgr.history)
	} else {
		// create a new conversation
		cid, _ := uuid.NewRandom()
		title := fmt.Sprintf("%s-%s", uid, cid.String())
		mgr.curCvs = &model.Conversation{
			ID:        cid,
			UserID:    uid,
			Title:     title,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			// Metadata:  map[string]any{},
		}
		if _, err := dao.AddConversation(context.Background(), mgr.curCvs); err != nil {
			return nil, fmt.Errorf("failed to add conversation: %w", err)
		}
	}
	return mgr, nil
}

type MemoryMgr struct {
	mux     sync.Mutex
	curCvs  *model.Conversation
	history []*schema.Message
}

func (memMgr *MemoryMgr) StoreMessageToDao(ctx context.Context, role schema.RoleType, message string) error {
	if message == "" {
		return fmt.Errorf("The message is empty")
	}
	if memMgr.curCvs == nil {
		return fmt.Errorf("The current conversation is not set")
	}
	if err := dao.AddMessage(ctx, (memMgr.curCvs.ID).String(), string(role), message); err != nil {
		return fmt.Errorf("failed to store message to dao: %w", err)
	}
	return nil
}

func (memMgr *MemoryMgr) GetCurConversation() (*model.Conversation, error) {
	return memMgr.curCvs, nil
}

func (memMgr *MemoryMgr) GetHistory() ([]*schema.Message, error) {
	return memMgr.history, nil
}

func (memMgr *MemoryMgr) SetCurConversation(conversation *model.Conversation) error {
	memMgr.curCvs = conversation
	return nil
}

func (memMgr *MemoryMgr) AppendMessage(ctx context.Context, role schema.RoleType, message string) error {
	memMgr.mux.Lock()
	defer memMgr.mux.Unlock()
	if message == "" {
		return fmt.Errorf("The message is empty")
	}
	if memMgr.curCvs == nil {
		return fmt.Errorf("The current conversation is not set")
	}
	memMgr.history = append(memMgr.history, &schema.Message{
		Role:    role,
		Content: message,
		// MultiContent: []schema.ChatMessagePart{},
		// Name:         "",
		// ToolCalls:    []schema.ToolCall{},
		// ToolCallID:   "",
		ResponseMeta: &schema.ResponseMeta{},
		Extra:        map[string]any{},
	})

	if err := memMgr.StoreMessageToDao(ctx, role, message); err != nil {
		return fmt.Errorf("failed to append message: %w", err)
	}

	return nil
}
