package initialize

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/document/parser"
	"github.com/cloudwego/eino/schema"
)

func NewLocalFileLoader(ctx context.Context) (*file.FileLoader, error) {
	loader, err := file.NewFileLoader(ctx, &file.FileLoaderConfig{
		UseNameAsID: true,                // Whether to use the file name as the document ID
		Parser:      &parser.ExtParser{}, // Optional: specify a custom parser
	})
	if err != nil {
		return nil, err
	}
	return loader, nil
}

func LoadLocalFiles(ctx context.Context, URI string) ([]*schema.Document, error) {
	loader, err := NewLocalFileLoader(ctx)
	if err != nil {
		return nil, err
	}

	documents, err := loader.Load(ctx, document.Source{
		URI: URI,
	})
	if err != nil {
		return nil, err
	}

	return documents, nil
}
