package main

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/erasernoob/JARVIS/auth"
	g "github.com/erasernoob/JARVIS/global"
	"github.com/erasernoob/JARVIS/graph/ragagent"
	"github.com/erasernoob/JARVIS/initialize"
	"github.com/erasernoob/JARVIS/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			log.Fatalf("panic occurred: %v", r)
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

	// agent := g.Agent
	// content, _ := service.SendUserMessage(ctx, agent, "my name is earsernoob")
	// fmt.Println(content)
	// content, _ = service.SendUserMessage(ctx, agent, "what's my name? and tell me my chat history use markdown")
	// fmt.Println(content)

	reader, err := ragagent.RunTheRagAgent(ctx, "请告诉我关于 Eino 的信息，越详细越好。")
	if err != nil {
		log.Fatalf("run rag agent failed: %s", err)
	}
	utils.StreamPrint(reader)
}
