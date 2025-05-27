package agent

// import (
// 	"context"
// 	"fmt"

// 	"github.com/erasernoob/JARVIS/beans"
// 	"github.com/erasernoob/JARVIS/cons"
// 	// db "github.com/erasernoob/JARVIS/memory"
// )

// func RestoreAgentFromDB(ctx context.Context) (*beans.Client, error) {
// 	var agent beans.Client
// 	// 查看是否有对话历史
// 	uid := ctx.Value(cons.UID).(string)

// 	// conversations, err := db.GetConversationByUid(ctx, uid)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get conversations: %w", err)
// 	}
// 	if len(conversations) != 0 {

// 	}

// 	// 如果有对话历史，保存后返回

// 	// 如果没有对话历史，创建一个新的对话
// }
