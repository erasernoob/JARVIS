package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid" // 用于处理 UUID
)

// Conversation 对应 conversations 表
type Conversation struct {
	ID        uuid.UUID       `json:"id"`
	UserID    string          `json:"user_id"` // sql.NullString 处理可空的 TEXT 类型
	Title     string          `json:"title"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Metadata  json.RawMessage `json:"metadata"` // json.RawMessage 处理 JSONB 类型
}

// 对应messages表
type Message struct {
	ID             uuid.UUID `json:"id"`
	ConversationID uuid.UUID `json:"conversation_id"`
	Role           string    `json:"role"`
	Content        string    `json:"content"`
	Timestamp      time.Time `json:"timestamp"`
	// SequenceNum    int             `json:"sequence_num"`
	Metadata json.RawMessage `json:"metadata"`
}
