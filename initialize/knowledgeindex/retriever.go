package knowledgeindex

import (
	"context"

	"github.com/cloudwego/eino-ext/components/retriever/redis"
	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/initialize/db"
)

func newRedisRetriever(ctx context.Context) (rtr retriever.Retriever, err error) {
	redisCli := db.NewRedisClient(ctx)

	config := &redis.RetrieverConfig{
		Client:            &redisCli,
		Index:             "",
		VectorField:       "",
		DistanceThreshold: new(float64),
		Dialect:           0,
		ReturnFields:      []string{},
		DocumentConverter: func(ctx context.Context, doc .Document) (*schema.Document, error) {
			panic("TODO")
		},
		TopK:      0,
		Embedding: nil,
	}

	return

}

type Reader interface {
	Read() string
}

type MyReader struct {
}
