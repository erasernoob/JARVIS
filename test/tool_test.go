package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"

	"github.com/erasernoob/JARVIS/global"
	mt "github.com/erasernoob/JARVIS/tool"
)

func Test_tool(t *testing.T) {
	InitTestEnv()
	fmt.Println(123)

	ctx := context.Background()

	// 1. Initialize tools
	searchTool := mt.GetDuckDuckGoSearchTool(ctx)
	todoTools := []tool.BaseTool{
		// getAddTodoTool(), // NewTool construction
		// updateTool,       // InferTool construction
		// &ListTodoTool{},  // Implements Tool interface
		searchTool, // Officially packaged tool
	}

	chatModel := global.Agent.LLM

	// 2. Get tool information and bind to ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(todoTools))
	for _, tool := range todoTools {
		info, err := tool.Info(ctx)
		if err != nil {
			log.Fatal(err)
		}
		toolInfos = append(toolInfos, info)
	}
	var err error
	chatModel, err = chatModel.WithTools(toolInfos)
	if err != nil {
		log.Fatal(err)
	}

	// 3.Create tools node
	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 4. Build the complete processing chain
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	// Compile and run the chain
	agent, err := chain.Compile(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Run example
	resp, err := agent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "Add a TODO for learning Eino and search for the repository address of cloudwego/eino",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output the result
	for _, msg := range resp {
		fmt.Println(msg.Content)
	}
}
