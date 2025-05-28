package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudwego/eino/components/document"
	"github.com/erasernoob/JARVIS/graph/knowledgeindexing"
)

func TestIndexing(t *testing.T) {
	InitTestEnv()
	ctx := context.Background()

	runner, err := knowledgeindexing.BuildKnowledgeIndexing(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := runner.Invoke(ctx, document.Source{
		URI: "./README.md",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
