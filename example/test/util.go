package otherpackage

import (
	"context"

	"github.com/rendis/abslog/v3"
)

func PrintFromOtherPackage() {
	var ctx = context.Background()
	ctxValuesMap := map[string]any{
		"id":   "1234567",
		"name": "John Doe",
		"age":  30,
	}
	ctx = context.WithValue(ctx, abslog.GetCtxKey(), ctxValuesMap)

	abslog.DebugCtx(ctx, "Default (Zap) Debug with context from other package")
	abslog.InfoCtx(ctx, "Default (Zap) Info with context from other package")
	abslog.WarnCtx(ctx, "Default (Zap) Warn with context from other package")
	abslog.ErrorCtx(ctx, "Default (Zap) Error with context from other package")
}
