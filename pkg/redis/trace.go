package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/extra/rediscmd/v8"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type TraceHook struct {
	tracer trace.Tracer
}

func NewTraceHook(tracer trace.Tracer) *TraceHook {
	return &TraceHook{tracer: tracer}
}

func (h TraceHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	ctx, span := h.tracer.Start(ctx, fmt.Sprintf("redis.process.%s", cmd.FullName()))
	span.SetAttributes(
		attribute.String("redis.cmd", rediscmd.CmdString(cmd)),
	)
	return ctx, nil
}

func (TraceHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	var errMsg string
	var isNil bool
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			isNil = true
		} else {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			errMsg = err.Error()
		}
	}

	span.SetAttributes(
		attribute.Bool("redis.nil", isNil),
		attribute.String("redis.error", errMsg),
	)

	return nil
}

func (h TraceHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	summary, cs := rediscmd.CmdsString(cmds)
	ctx, span := h.tracer.Start(ctx, fmt.Sprintf("redis.process.pipeline:%s", summary))
	span.SetAttributes(
		attribute.String("redis.cmd", cs),
	)
	return ctx, nil
}

func (TraceHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	var errs []string
	var errMsg string

	for _, cmd := range cmds {
		if err := cmd.Err(); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		errMsg = errs[0]
		span.RecordError(errors.New(errMsg))
		span.SetStatus(codes.Error, errMsg)
	}

	span.SetAttributes(
		attribute.String("redis.error", errMsg),
		attribute.StringSlice("redis.errors", errs),
	)

	return nil
}
