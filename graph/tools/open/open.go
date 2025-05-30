package open

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

// This is the open tool open file or url

type OpenURIImpl struct {
	config *OpenURIConfig
}

type OpenURIConfig struct{}

type OpenURIReq struct {
	URI string `json:"uri" jsonschema:"required description=The URI of the request resource(file of website)"`
}

func defaultOpenURIConfig(ctx context.Context) (*OpenURIConfig, error) {
	config := &OpenURIConfig{}
	return config, nil
}

func NewOpenFileTool(ctx context.Context, config *OpenURIConfig) (tn tool.BaseTool, err error) {
	if config == nil {
		config, err = defaultOpenURIConfig(ctx)
		if err != nil {
			return nil, err
		}
	}
	t := &OpenURIImpl{config: config}
	tn, err = t.ToEinoTool(ctx)
	if err != nil {
		return nil, err
	}
	return tn, nil
}

type OpenURIRes struct {
	Message string `json:"message" jsonschema:"description=The specific message of this operation"`
}

func (impl *OpenURIImpl) ToEinoTool(ctx context.Context) (t tool.BaseTool, err error) {
	return utils.InferTool("OpenURI", "Open the resource infer by the URI(file or the website)", impl.Invoke)
}

func (impl *OpenURIImpl) Invoke(ctx context.Context, req OpenURIReq) (res OpenURIRes, err error) {
	uri := req.URI
	if uri == "" {
		res.Message = "The uri should not be empty"
		return res, nil
	}

	// consider if it is a file path
	if isFilePath(req.URI) {
		req.URI = strings.TrimPrefix(req.URI, "file:///")
		if _, err := os.Stat(req.URI); err != nil {
			res.Message = fmt.Sprintf("file not exists: %s", req.URI)
			return res, nil
		}
	}

	err = openURI(req.URI)
	if err != nil {
		res.Message = fmt.Sprintf("failed to open the file: %s", req.URI)
		return res, nil
	}

	res.Message = fmt.Sprintf("success, open %s", req.URI)
	return res, nil
}

func openURI(uri string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", uri)
	case "darwin":
		cmd = exec.Command("open", uri)
	case "linux":
		cmd = exec.Command("xdg-open", uri)
	default:
		return fmt.Errorf("Unsupported Platform")
	}
	return cmd.Run()
}

func isFilePath(path string) bool {
	s, err := url.Parse(path)
	return err == nil && s.Scheme == "file" && s.Path != ""
}
