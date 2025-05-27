package service

import (
	"context"
	"fmt"

	"github.com/erasernoob/JARVIS/common"
	db "github.com/erasernoob/JARVIS/dao"
	g "github.com/erasernoob/JARVIS/global"
	m "github.com/erasernoob/JARVIS/model"
)

func RestoreClientFromDB(ctx context.Context) (*m.Client, error) {
	var agent m.Client
	// 查看是否有对话历史
	uid := ctx.Value(common.UID).(string)

	conversations, err := db.GetConversationByUid(ctx, g.PgConn, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversations: %w", err)
	}
	if len(conversations) != 0 {

	}

	// 如果有对话历史，保存后返回

	// 如果没有对话历史，创建一个新的对话
}
