package main

import (
	"context"
	"fmt"
	"log"

	"github.com/erasernoob/JARVIS/auth"
	g "github.com/erasernoob/JARVIS/global"
	"github.com/erasernoob/JARVIS/initialize"
)

func main() {
	ctx := context.Background()
	// mock the userID
	ctx = auth.Identify(ctx)

	if err := g.Init(ctx); err != nil {
		log.Fatalf("init failed: %s", err)
	}
	if err := initialize.Init(ctx); err != nil {
		log.Fatalf("initialize failed: %s", err)
	}

	agent := g.Agent
	content, _ := agent.SendUserMessage(ctx, "Tell me a joke, and it's about the programming")
	fmt.Println(content)
}
