package knowledge_index

import (
	"context"

	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/compose"
	"github.com/erasernoob/JARVIS/initialize"
)

const (
	FileLoader             = "fileLoader"
	MarkdownHeaderSplitter = "mdHeaderSplitter"
	RedisIndexer           = "RedisIndex"
	GraphName              = "KnowledgeIndexing"
)

// Runnable [Input, Output]
func BuildKnowledgeIndexing(ctx context.Context) (r compose.Runnable[document.Source, []string], err error) {
	// 1. create a new graph
	g := compose.NewGraph[document.Source, []string]()

	// 2. create a new file loader
	loader, err := initialize.NewLocalFileLoader(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddLoaderNode(FileLoader, loader)

	// 3. create a new transformer(splitter)
	mdHeaderSplitter, err := initialize.NewMdHeaderTransformer(ctx)
	if err != nil {
		return nil, err
	}
	- = g.AddDocumentTransformerNode(MarkdownHeaderSplitter, mdHeaderSplitter)

	// 4. create a new indexer
	redisIndexer, err := initialize.NewRedisIndexer(ctx)
	if err != nil {
		return nil, err
	}

	_ = g.AddIndexerNode(RedisIndexer, redisIndexer)

	// 5. create start to add branch

	_ = g.AddEdge(compose.START, FileLoader)
	_ = g.AddEdge(FileLoader, MarkdownHeaderSplitter)
	_ = g.AddEdge(MarkdownHeaderSplitter, RedisIndexer)
	_ = g.AddEdge(RedisIndexer, compose.END)

	r, err = g.Compile(ctx, compose.WithGraphName(GraphName), compose.WithNodeTriggerMode(compose.AnyPredecessor))
	if err != nil {
		return nil, err
	}
	return r, err
}
