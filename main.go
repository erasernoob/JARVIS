package main

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/erasernoob/JARVIS/auth"
	g "github.com/erasernoob/JARVIS/global"
	"github.com/erasernoob/JARVIS/initialize"
	"github.com/erasernoob/JARVIS/service"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic occurred: %v", r)
			debug.PrintStack()
		}
	}()
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
	content, _ := service.SendUserMessage(ctx, agent, "my name is earsernoob")
	fmt.Println(content)
	content, _ = service.SendUserMessage(ctx, agent, "what's my name? and tell me my chat history use markdown")
	fmt.Println(content)
}
