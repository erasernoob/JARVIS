package ragagent

import (
	"context"
	"time"
)

// create the inputToQuery Lambda node
func NewInputToQueryNode(ctx context.Context, input *UserMessage, opts ...any) (output string, err error) {
	// return only the query
	return input.Query, nil
}

func NewInputToHistoryNode(ctx context.Context, input *UserMessage, opts ...any) (output map[string]interface{}, err error) {
	return map[string]interface{}{
		"history": input.History,
		"content": input.Query,
		"date":    time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
