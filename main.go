package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

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
	RunRagAgentDialog("1")
}

func RunRagAgentDialog(uid string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("🧑‍: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		in := strings.TrimSpace(input)
		if in == "exit" || in == "quit" || in == " " {
			return
		}

		sr, err := ragagent.RunTheRagAgent(context.Background(), uid, input)
		if err != nil {
			fmt.Printf("Error running RAG agent: %s\n", err)
			continue
		}
		utils.StreamPrint(sr)
	}
}
