package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/eino/schema"
	g "github.com/erasernoob/JARVIS/global"
	"github.com/joho/godotenv"
)

func TestMain(t *testing.T) {
	_ = godotenv.Load()
	ctx := context.Background()
	_ = g.Init(ctx)

	content, _ := g.LLM.Generate(ctx, []*schema.Message{
		schema.SystemMessage("You are a very helpful assistant"),
	})
	fmt.Println(content)
}
