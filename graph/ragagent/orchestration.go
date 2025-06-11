package ragagent

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/graph/baseagent"
	"github.com/erasernoob/JARVIS/graph/ragagent/mem"
	"github.com/erasernoob/JARVIS/initialize/knowledgeindex"
	"google.golang.org/appengine/log"
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

var MemMgr mem.Memory

func Init(uid string) {
	// Initialize the memory manager
	var err error
	MemMgr, err = mem.NewMemoryMgr(uid)
	if err != nil {
		log.Errorf(context.Background(), "failed to initialize memory manager: %v", err)
		return
	}
}

func RunTheRagAgent(ctx context.Context, uid string, msg string) (output *schema.StreamReader[*schema.Message], err error) {
	Init(uid)

	// 1. build the rag agent
	ragAgent, err := BuildRagAgent(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to build rag agent: %w", err)
	}
	// Get the history from the database
	history, err := MemMgr.GetHistory()
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %w", err)
	}

	// 2. create the input message
	input := &UserMessage{
		ID:      uid,
		Query:   msg,
		History: history,
	}

	// 3. run the rag agent
	output, err = ragAgent.Stream(ctx, input)
	if err != nil {
		return nil, err
	}

	srs := output.Copy(2)

	// Save the Assistant's response to the memory
	go SaveStreamResponse(ctx, srs[1])

	err = MemMgr.AppendMessage(ctx, schema.Assistant, msg)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func SaveStreamResponse(ctx context.Context, output *schema.StreamReader[*schema.Message]) {
	// for save to memory
	fullMsgs := make([]*schema.Message, 0)

	defer func() {
		// close stream if you used it
		output.Close()

		// add user input to history
		fullMsg, err := schema.ConcatMessages(fullMsgs)
		if err != nil {
			fmt.Println("error concatenating messages: ", err.Error())
		}
		// add agent response to history
		if err := MemMgr.AppendMessage(ctx, schema.Assistant, fullMsg.Content); err != nil {
			log.Errorf(ctx, "error appending message to memory: %v\n", err.Error())
			return
		}

	}()

outer:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done", ctx.Err())
			return
		default:
			chunk, err := output.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break outer
				}
			}

			fullMsgs = append(fullMsgs, chunk)
		}
	}
}
