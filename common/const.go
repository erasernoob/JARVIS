package common

type contextKey string

// To avoid the context key collision, we define a custom type for the context key.
// This is a unique key to store the user ID in the context.
// SA1029

const (
	UID contextKey = "user_id"
)

const (
	ASSISTANT = "assistant"
	USER      = "user"
)
