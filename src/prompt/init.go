package prompt

import (
	"log"
	"os"

	"github.com/cloudwego/eino/schema"
)

var (
	BASE_PROMPT = getBasePrompt()
)

func getBasePrompt() *schema.Message {
	bytes, err := os.ReadFile("./src/prompt/base.md")
	if err != nil {
		log.Fatalf("read the base prompt failed: %s", err)
	}
	content := string(bytes)
	return schema.SystemMessage(content)

}
