package main

import (
	"context"
	"fmt"
	"log"

	g "github.com/erasernoob/agent/src/global"
)

func main() {
	ctx := context.Background()
	if err := g.Init(ctx); err != nil {
		log.Fatalf("init failed: %s", err)
	}
	agent := g.GetAgent()
	content, _ := agent.SendUserMessage(ctx, "Tell me a joke, and it's about the programming")
	fmt.Println(content)
}
