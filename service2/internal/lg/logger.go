package lg

import (
	"context"
	"service2/internal/handler/middlewares"

	"github.com/pkg/errors"
	"github.com/vrnvgasu/logwrapper"
)

var Logger *logwrapper.StandardLogger

func Error(ctx context.Context, op string, pack string, err error) {
	Logger.Payload(newPayload(ctx, op, pack)).Error(errors.WithStack(err))
}

func Fatal(ctx context.Context, op string, pack string, err error) {
	Logger.Payload(newPayload(ctx, op, pack)).Fatal(errors.WithStack(err))
}

func Info(ctx context.Context, op string, pack string, message string) {
	Logger.Payload(newPayload(ctx, op, pack)).Info(message)
}

func newPayload(ctx context.Context, op string, pack string) *logwrapper.Payload {
	p := logwrapper.NewPayload().
		Op(op).
		Package(pack)

	switch ctx.Value(middlewares.RequestID).(type) {
	case string:
		p.CtxID(ctx, middlewares.RequestID)
	}

	return p
}
