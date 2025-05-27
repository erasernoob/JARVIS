package auth

import (
	"context"

	"github.com/erasernoob/JARVIS/cons"
)

func GetCurUserID() string {
	return "123"
}

func Identify(ctx context.Context) context.Context {
	uid := GetCurUserID()
	ctxWithValue := context.WithValue(ctx, cons.UID, uid)
	return ctxWithValue
}

func init() {

}
