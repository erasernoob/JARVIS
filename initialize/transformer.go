package initialize

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
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
