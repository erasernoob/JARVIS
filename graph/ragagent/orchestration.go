package ragagent

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/graph/baseagent"
	"github.com/erasernoob/JARVIS/initialize/knowledgeindex"
)

func BuildRagAgent(ctx context.Context) (r compose.Runnable[*UserMessage, *schema.Message], err error) {
	const (
		// InputToQuery   = "InputToQuery"
		ChatTemplate   = "ChatTemplate"
		ReactAgent     = "ReactAgent"
		RedisRetriever = "RedisRetriever"
		InputToHistory = "InputToHistory"
	)

	// 1. var the graph
	graph := compose.NewGraph[*UserMessage, *schema.Message]()

	// 2. new the agent node
	agent, err := baseagent.BuildBaseAgent(ctx)
	if err != nil {
		return nil, err
	}
	// graph.AddLambdaNode(ReactAgent, agent, )

	// 3. new the retriever node
	retriever, err := knowledgeindex.NewRedisRetriever(ctx)
	if err != nil {
		return nil, err
	}
	// 4. create a the chatTemplate
	chatTemplate, err := NewChatTemplate(ctx)
	if err != nil {
		return nil, err
	}

	// 5. compile the graph

	return
}
