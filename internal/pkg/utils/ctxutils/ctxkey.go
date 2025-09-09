package ctxutils

type CtxKey string

const (
	keyTraceID CtxKey = "trace_id"
	keyUserID  CtxKey = "user_id"
)
