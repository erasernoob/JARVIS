package knowledgeindex

import (
	"context"
	"fmt"
	"strconv"

	redisClient "github.com/redis/go-redis/v9"

	"github.com/cloudwego/eino-ext/components/retriever/redis"
	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/initialize/db"
)

func NewRedisRetriever(ctx context.Context) (rtr retriever.Retriever, err error) {
	redisCli := db.NewRedisClient(ctx)

	config := &redis.RetrieverConfig{
		Client:      &redisCli,
		Index:       fmt.Sprintf("%s%s", RedisPrefix, IndexName),
		VectorField: VectorField,
		// DistanceThreshold: new(float64),
		Dialect: 2,
		// 需要返回的Fields
		ReturnFields: []string{ContentField, MetadataField, VectorField},
		// 将redis中的文档转换成schema.Document
		DocumentConverter: func(ctx context.Context, doc redisClient.Document) (*schema.Document, error) {
			resp := &schema.Document{
				ID:       doc.ID,
				Content:  "",
				MetaData: map[string]any{},
			}
			for field, val := range doc.Fields {
				if field == ContentField {
					resp.Content = val
				} else if field == MetadataField {
					resp.MetaData[field] = val
				} else if field == DistanceField {
					distance, err := strconv.ParseFloat(val, 64)
					if err != nil {
						continue
					}
					resp.WithScore(1 - distance)
				}
			}
			return resp, nil
		},
		TopK: 8,
		// Embedding: (ctx),
	}
	embedding, err := newEmbedding(ctx)
	if err != nil {
		return nil, err
	}
	config.Embedding = embedding
	rtr, err = redis.NewRetriever(ctx, config)
	if err != nil {
		return nil, err
	}
	return
}
