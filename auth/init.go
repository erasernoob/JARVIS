package auth

import (
	"context"

	cons "github.com/erasernoob/JARVIS/common"
)

func GetCurUserID() string {
	return "1234"
}

func Identify(ctx context.Context) context.Context {
	uid := GetCurUserID()
	ctxWithValue := context.WithValue(ctx, cons.UID, uid)
	return ctxWithValue
}

func init() {

}
