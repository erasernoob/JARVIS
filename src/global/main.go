package global

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/agent/src/prompt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const (
	MODEL_NAME = "MODEL_NAME"
	API_KEY    = "DASHSCOPE_API_KEY"
	BASE_URL   = "OPENAI_BASE_URL"
)

var (
	LLM     model.ToolCallingChatModel
	History []*schema.Message
	Agent   *Client
)

func Init(ctx context.Context) error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	var err error
	LLM, err = createChatModel(ctx)
	if err != nil {
		return err
	}

	// Get the Base Prompt
	if err = getBasePrompt(ctx); err != nil {
		return err
	}

	// initialize the agent
	if err = initializeAgent(ctx); err != nil {
		return err
	}

	return nil
}

func GetAgent() *Client {
	return Agent
}

func getBasePrompt(c context.Context) error {
	History = append(History,
		prompt.BASE_PROMPT,
	)
	return nil
}

func initializeAgent(ctx context.Context) error {
	Agent = &Client{
		LLM:     LLM,
		History: History,
	}
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
