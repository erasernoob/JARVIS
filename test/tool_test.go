package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/global"
	"github.com/erasernoob/JARVIS/graph/tools/open"
)

func TestTool(t *testing.T) {
	InitTestEnv()
	ctx := context.Background()

	opentool, err := open.NewOpenFileTool(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	openInfo, err := opentool.Info(ctx)
	if err != nil {
		fmt.Println(err)
	}
	LLM, err := global.Agent.LLM.WithTools([]*schema.ToolInfo{openInfo})
	if err != nil {
		fmt.Println(err)
	}

	toolArray := []tool.BaseTool{
		opentool,
	}

	// start to create a chain
	global.Agent.LLM = LLM
	// content, err := service.SendUserMessage(ctx, global.Agent, "Help me to open the file in my 'D' disk named mapreduce.pdf")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()

	toolNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: toolArray,
	})
	if err != nil {
		fmt.Println(err)
	}

	chain.AppendChatModel(global.Agent.LLM)
	chain.AppendToolsNode(toolNode)
	chain.AppendChatModel(global.Agent.LLM)

	c, err := chain.Compile(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := c.Invoke(ctx, []*schema.Message{
		schema.SystemMessage("You are a very helpful agent"),
		schema.UserMessage("Help me to open the file in my 'D' disk named mapreduce.pdf"),
	})

	if err != nil {
		fmt.Println(err)
	}

	for _, msg := range res {
		fmt.Println(msg.Content)
	}
}

func TestToolUseRecAgent(t *testing.T) {
	InitTestEnv()
	ctx := context.Background()

	opentool, err := open.NewOpenFileTool(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	openInfo, err := opentool.Info(ctx)
	if err != nil {
		fmt.Println(err)
	}
	LLM, err := global.Agent.LLM.WithTools([]*schema.ToolInfo{openInfo})
	if err != nil {
		fmt.Println(err)
	}

	toolArray := []tool.BaseTool{
		opentool,
	}

	agent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: LLM,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: toolArray,
		},
	})

	res, err := agent.Generate(ctx, []*schema.Message{
		schema.SystemMessage("You are a very helpful agent"),
		schema.UserMessage("help me open the 'd:/documents/mapreduce.pdf' file give all of your tools call information make me debug easily"),
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Content)

}
