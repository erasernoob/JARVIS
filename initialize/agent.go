package initialize

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"

	m "github.com/erasernoob/JARVIS/model"
	"github.com/erasernoob/JARVIS/service"

	"github.com/erasernoob/JARVIS/prompt"
)

const (
	MODEL_NAME = "MODEL_NAME"
	API_KEY    = "DASHSCOPE_API_KEY"
	BASE_URL   = "OPENAI_BASE_URL"
)

var (
	LLM     model.ToolCallingChatModel
	History []*schema.Message
	Agent   *m.Client
)

func init() {

}

func InitAgent(ctx context.Context) (*m.Client, error) {
	if Agent != nil {
		return Agent, nil
	}

	if err := getBasePrompt(ctx); err != nil {
		log.Fatalf("Get base prompt failed: err=%v", err)
		return nil, err
	}

	chatModel, err := createChatModel(ctx)
	if err != nil {
		log.Fatalf("Create chat model failed: err=%v", err)
		return nil, err
	}
	LLM = chatModel

	if err := initializeAgent(ctx); err != nil {
		log.Fatalf("Initialize agent failed: err=%v", err)
		return nil, err
	}
	return Agent, nil
}

func getBasePrompt(c context.Context) error {
	History = append(History,
		prompt.BASE_PROMPT,
	)
	return nil
}

func initializeAgent(ctx context.Context) error {
	// Get the memory restore the history from database

	agent, err := service.RestoreClientFromDB(ctx)
	if err != nil {
		log.Fatalf("Restore client from DB failed: err=%v", err)
		return err
	}
	agent.LLM = LLM
	Agent = agent
	// 叠加
	Agent.History = append(History, agent.History...)
	return nil
}

func createChatModel(ctx context.Context) (model.ToolCallingChatModel, error) {
	key := os.Getenv(API_KEY)
	modelName := os.Getenv(MODEL_NAME)
	baseURL := os.Getenv(BASE_URL)

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: baseURL,
		Model:   modelName,
		APIKey:  key,
	})
	if err != nil {
		log.Fatalf("Create the ChatModel failed: err=%v", err)
	}
	return chatModel, nil
}
