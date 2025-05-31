package baseagent

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/erasernoob/JARVIS/global"
)

func GetTools(ctx context.Context) (tools []tool.BaseTool, err error) {
	return tools, err
}

// use the react Agent provided by the eino
func BuildBaseAgent(ctx context.Context) (lba *compose.Lambda, err error) {
	config := react.AgentConfig{
		MaxStep:          25,
		ToolCallingModel: global.LLM,
		// add or remove the prompt msg during the message transport
		// MessageModifier: ,
		ToolReturnDirectly: map[string]struct{}{},
	}

	tools, err := GetTools(ctx)
	if err != nil {
		return nil, err
	}

	config.ToolsConfig.Tools = tools

	ag, err := react.NewAgent(ctx, &config)
	if err != nil {
		return nil, err
	}
	// create the lambda in the orchestration

	lba, err = compose.AnyLambda(ag.Generate, ag.Stream, nil, nil)
	if err != nil {
		return nil, err
	}
	return
}
