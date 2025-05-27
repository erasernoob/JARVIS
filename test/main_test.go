package test

import (
	"context"
	"testing"

	g "github.com/erasernoob/JARVIS/global"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.T) {
	_ = godotenv.Load()
	ctx := context.Background()
	_ = g.Init(ctx)

	// Test_tool()

}
