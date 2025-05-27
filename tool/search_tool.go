package tool

import (
	"context"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
)

func GetDuckDuckGoSearchTool(ctx context.Context) tool.InvokableTool {
	duckduck, _ := duckduckgo.NewTool(ctx, &duckduckgo.Config{})
	return duckduck
}
