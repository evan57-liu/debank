package ctxutils

import (
	"context"
	"time"
)

func GetTraceID(ctx context.Context) string {
	traceID, _ := ctx.Value(keyTraceID).(string)

	return traceID
}

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, keyTraceID, traceID)
}

func GetUserID(ctx context.Context) string {
	traceID, _ := ctx.Value(keyUserID).(string)

	return traceID
}

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, keyUserID, userID)
}

// Detach returns a context that keeps all the values of its parent context
// but detaches from the cancellation and error handling.
func Detach(ctx context.Context) context.Context {
	return detachedContext{parent: ctx}
}

type detachedContext struct {
	parent context.Context
}

func (detachedContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (detachedContext) Done() <-chan struct{} {
	return nil
}

func (detachedContext) Err() error {
	return nil
}

func (d detachedContext) Value(key interface{}) interface{} {
	return d.parent.Value(key)
}
