package knowledgeindex

import (
	"context"
	"encoding/json"
	"fmt"

	// "github.com/cloudwego/eino-examples/quickstart/eino_assistant/pkg/redis"
	"github.com/cloudwego/eino-ext/components/indexer/redis"
	"github.com/erasernoob/JARVIS/initialize/db"
	"github.com/google/uuid"

	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/schema"
	// redisCli "github.com/redis/go-redis/v9"
)

const (
	RedisPrefix = "JARVIS:doc:"
	IndexName   = "vector_index"

	ContentField  = "content"
	MetadataField = "metadata"
	VectorField   = "content_vector"
	DistanceField = "distance"
)

// Use the redis as the vector store(但其实感觉更像是存储分块完之后的文本的地方)
// indexer 包括了embedding模型
func NewRedisIndexer(ctx context.Context) (idr indexer.Indexer, err error) {
	redisCli := db.NewRedisClient(ctx)

	// start the build the redis.config

	config := &redis.IndexerConfig{
		Client:    &redisCli,
		KeyPrefix: RedisPrefix,
		BatchSize: 1,
		// 将文件转变成`hash`数据结构
		DocumentToHashes: func(ctx context.Context, doc *schema.Document) (*redis.Hashes, error) {
			if doc.ID == "" {
				doc.ID = uuid.New().String()
			}
			key := doc.ID

			metadataBytes, err := json.Marshal(doc.MetaData)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal metadata: %w", err)
			}

			return &redis.Hashes{
				Key: key,
				Field2Value: map[string]redis.FieldValue{
					ContentField:  {Value: doc.Content, EmbedKey: VectorField},
					MetadataField: {Value: metadataBytes},
				},
			}, nil
		},
	}
	// create a new embedding
	embedding, err := newEmbedding(ctx)
	if err != nil {
		return nil, err
	}

	config.Embedding = embedding

	idr, err = redis.NewIndexer(ctx, config)
	if err != nil {
		return nil, err
	}
	return idr, nil
}
