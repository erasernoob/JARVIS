package initialize

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/erasernoob/JARVIS/common"
)

func newEmbedding(ctx context.Context) (eb embedding.Embedder, err error) {
	config := &openai.EmbeddingConfig{
		BaseURL: os.Getenv(common.BASE_URL),
		APIKey:  os.Getenv(common.API_KEY),
		Model:   os.Getenv(common.EMBEDDING_MODEL),
	}

	// create new embedder
	eb, err = openai.NewEmbedder(ctx, config)
	if err != nil {
		return nil, err
	}
	return eb, nil
}
