package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/erasernoob/JARVIS/initialize/knowledgeindex"
)

func TestIndexing(t *testing.T) {
	InitTestEnv()
	ctx := context.Background()

	rtr, _ := knowledgeindex.NewRedisRetriever(ctx)
	docs, _ := rtr.Retrieve(ctx, "抽象工厂模式")
	fmt.Println("Retrieved documents:", len(docs))
	log.Printf("vikingDB retrieve success, query=%v, docs=%v", "", docs)

	// runner, err := knowledgeindexing.BuildKnowledgeIndexing(ctx)
	// if err != nil {
	// 	debug.Stack()
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// res, err := runner.Invoke(ctx, document.Source{
	// 	// URI: "./files/abstract-factory.md",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(res)
}
