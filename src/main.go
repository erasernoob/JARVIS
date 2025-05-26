package main

import (
	"context"
	"fmt"
	"log"

	g "github.com/erasernoob/JARVIS/src/global"
)

func main() {
	ctx := context.Background()
	if err := g.Init(ctx); err != nil {
		log.Fatalf("init failed: %s", err)
	}
	agent := g.Agent
	content, _ := agent.SendUserMessage(ctx, "Tell me a joke, and it's about the programming")
	fmt.Println(content)
}
