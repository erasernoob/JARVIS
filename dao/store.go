package dao

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/schema"
	g "github.com/erasernoob/JARVIS/global"
	"github.com/erasernoob/JARVIS/model"
)

func GetMessagesByConversationID(ctx context.Context, cid string) ([]*schema.Message, error) {
	var res []*schema.Message

	sql := `SELECT * FROM messages WHERE conversation_id = $1 ORDER BY created_at DESC`
	rows, err := g.PgConn.Query(ctx, sql, cid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		// 声明值类型，分配内存
		var msg model.Message
		_ = rows.Scan(
			&msg.ID,
			&msg.ConversationID,
			&msg.Role,
			&msg.Content,
			&msg.Timestamp,
			&msg.Metadata,
		)
		res = append(res, &schema.Message{
			Role:    schema.RoleType(msg.Role),
			Content: msg.Content,
			// Timestamp: msg.CreatedAt,
			// Metadata:  msg.Metadata,
		})
	}
	return res, nil
}

func GetConversationByUid(ctx context.Context, uid string) ([]*model.Conversation, error) {
	var res []*model.Conversation

	sql := `SELECT * FROM conversations WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := g.PgConn.Query(ctx, sql, uid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		// 声明值类型，分配内存
		var conv model.Conversation
		_ = rows.Scan(
			&conv.ID,
			&conv.UserID,
			&conv.Title,
			&conv.CreatedAt,
			&conv.UpdatedAt,
			&conv.Metadata,
		)
		res = append(res, &conv)
	}
	return res, nil
}

func AddMessage(ctx context.Context, cid string, role string, content string) error {
	sql := `INSERT INTO messages (conversation_id, role, content, timestamp) VALUES ($1, $2, $3, NOW())`
	t, err := g.PgConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}

	tag, err := t.Exec(ctx, sql, cid, role, content)
	if err != nil || tag.RowsAffected() == 0 {
		// 回滚事务
		if rbErr := t.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("rollback failed: %w, original error: %v", rbErr, err)
		}
		return fmt.Errorf("exec addUser Message failed: %w", err)
	}
	return nil
}

func AddConversation(ctx context.Context, c *model.Conversation) (string, error) {
	sql := `INSERT INTO conversations (id, user_id, title, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	var id string
	err := g.PgConn.QueryRow(ctx, sql, c.ID, c.UserID, c.Title).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("insert conversation failed: %w", err)
	}

	if err != nil {
		return "", fmt.Errorf("parse conversation ID failed: %w", err)
	}
	return id, nil
}
