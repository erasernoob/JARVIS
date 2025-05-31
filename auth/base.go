package auth

import (
	"context"

	cons "github.com/erasernoob/JARVIS/common"
)

func GetCurUserID() string {
	return "testtt"
}

func Identify(ctx context.Context) context.Context {
	uid := GetCurUserID()
	ctxWithValue := context.WithValue(ctx, cons.UID, uid)
	return ctxWithValue
}

func init() {

}
