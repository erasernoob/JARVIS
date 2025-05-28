package common

type contextKey string

// To avoid the context key collision, we define a custom type for the context key.
// This is a unique key to store the user ID in the context.
// SA1029

const (
	UID contextKey = "user_id"
)

const (
	ASSISTANT  = "assistant"
	USER       = "user"
	REDIS_ADDR = "REDIS_ADDR"
)

const (
	MODEL_NAME      = "MODEL_NAME"
	API_KEY         = "DASHSCOPE_API_KEY"
	BASE_URL        = "OPENAI_BASE_URL"
	EMBEDDING_MODEL = "EMBEDDING_MODEL"
)
