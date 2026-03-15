package trace

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type ctxKey string //to avoid collision with context value

const traceIDKey ctxKey = "trace_id"

func NewTraceID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func TraceIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(traceIDKey)
	s, ok := v.(string)
	return s, ok && s != ""
}
