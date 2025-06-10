package ragagent

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/graph/baseagent"
	"github.com/erasernoob/JARVIS/initialize/knowledgeindex"
)

func BuildRagAgent(ctx context.Context) (r compose.Runnable[*UserMessage, *schema.Message], err error) {
	const (
		InputToQuery   = "InputToQuery"
		ChatTemplate   = "ChatTemplate"
		ReactAgent     = "ReactAgent"
		RedisRetriever = "RedisRetriever"
		InputToHistory = "InputToHistory"

		// to adapt the system's prompt
		RetrieverOutputKey = "documents"
	)

	// 1. var the graph
	graph := compose.NewGraph[*UserMessage, *schema.Message]()
	// 2. new the agent node
	agent, err := baseagent.BuildBaseAgent(ctx)
	if err != nil {
		return nil, err
	}
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
	// 5. create the input lambda node
	input2queryLambda := compose.InvokableLambdaWithOption(NewInputToQueryNode)
	// 6. create the input2history lambda node
	input2historyLambda := compose.InvokableLambdaWithOption(NewInputToHistoryNode)

	// Add Node
	_ = graph.AddLambdaNode(ReactAgent, agent, compose.WithNodeName("React Agent"))
	_ = graph.AddRetrieverNode(RedisRetriever, retriever, compose.WithNodeName("Redis Retriever"), compose.WithOutputKey(RetrieverOutputKey))
	_ = graph.AddChatTemplateNode(ChatTemplate, chatTemplate, compose.WithNodeName("Chat Template"))

	_ = graph.AddLambdaNode(InputToQuery, input2queryLambda)
	_ = graph.AddLambdaNode(InputToHistory, input2historyLambda)

	// Add Edge
	_ = graph.AddEdge(compose.START, InputToQuery)
	_ = graph.AddEdge(compose.START, InputToHistory)

	// 同时完成两个前置条件后，编排继续往下走
	_ = graph.AddEdge(InputToQuery, RedisRetriever)
	_ = graph.AddEdge(InputToHistory, ChatTemplate)
	_ = graph.AddEdge(RedisRetriever, ChatTemplate)
	_ = graph.AddEdge(ChatTemplate, ReactAgent)
	_ = graph.AddEdge(ReactAgent, compose.END)
	// 5. compile the graph

	graphOption := []compose.GraphCompileOption{
		compose.WithGraphName("RAG Agent"),
		// 设置节点触发模式
		// all_predecessor: 只有当所有前置节点都完成后，当前节点才会被触发
		compose.WithNodeTriggerMode(compose.AllPredecessor),
	}

	r, err = graph.Compile(ctx, graphOption...)
	return
}

func RunTheRagAgent(ctx context.Context, msg string) (output *schema.StreamReader[*schema.Message], err error) {
	// 1. build the rag agent
	ragAgent, err := BuildRagAgent(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to build rag agent: %w", err)
	}

	// 2. create the input message
	input := &UserMessage{
		Query: msg,
	}

	// 3. run the rag agent
	output, err = ragAgent.Stream(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}
