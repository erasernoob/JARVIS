package knowledgeindex

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/recursive"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/semantic"
	"github.com/cloudwego/eino/components/document"
)

// Markdown
func NewMdHeaderTransformer(ctx context.Context) (dtf document.Transformer, err error) {
	// 这里的设置会影响到最后分词得到的metadata内容
	config := &markdown.HeaderConfig{
		Headers: map[string]string{
			"#": "Title",
		},
		TrimHeaders: true, // 是否去除标题前的空格
	}
	// New 根据header来进行分词
	dtf, err = markdown.NewHeaderSplitter(ctx, config)
	if err != nil {
		return nil, err
	}
	return
}

// splitter semantic
func NewSemanticTransformer(ctx context.Context) (tsf document.Transformer, err error) {
	embedder, err := newEmbedding(ctx)
	if err != nil {
		return nil, err
	}
	config := &semantic.Config{
		Embedding:    embedder,                      // 必需：用于生成文本向量的嵌入器
		BufferSize:   2,                             // 可选：上下文缓冲区大小
		MinChunkSize: 10,                            // 可选：最小片段大小
		Separators:   []string{"\n", ".", "?", "!"}, // 可选：分隔符列表
		Percentile:   0.9,                           // 可选：分割阈值百分位数
		LenFunc:      nil,                           // 可选：自定义长度计算函数

	}

	tsf, err = semantic.NewSplitter(ctx, config)
	return
}

// recursive splitter
func NewRecursiveTransformer(ctx context.Context) (tsf document.Transformer, err error) {
	config := &recursive.Config{
		ChunkSize:   1000,
		OverlapSize: 200,
		Separators:  []string{"\n\n", "\n", "。", "！", "？"},
		// keep the separators at the end
		KeepType: recursive.KeepTypeEnd,
	}

	tsf, err = recursive.NewSplitter(ctx, config)
	return
}
