package db

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	PluginTrace = "PLUGIN_TRACE"
	spanBegin   = "trace:span_begin"
	spanEnd     = "trace:span_end"
)

type TracePlugin struct {
	tracer trace.Tracer
}

func NewTracePlugin(tracer trace.Tracer) TracePlugin {
	return TracePlugin{tracer: tracer}
}

func (TracePlugin) Name() string {
	return PluginTrace
}

func (cp TracePlugin) Initialize(db *gorm.DB) error {

	_ = db.Callback().Create().Before("gorm:before_create").Register(spanBegin, cp.spanBegin("create"))
	_ = db.Callback().Create().After("gorm:after_create").Register(spanEnd, cp.spanEnd)

	_ = db.Callback().Query().Before("gorm:query").Register(spanBegin, cp.spanBegin("query"))
	_ = db.Callback().Query().After("gorm:after_query").Register(spanEnd, cp.spanEnd)

	_ = db.Callback().Delete().Before("gorm:before_delete").Register(spanBegin, cp.spanBegin("delete"))
	_ = db.Callback().Delete().After("gorm:after_delete").Register(spanEnd, cp.spanEnd)

	_ = db.Callback().Update().Before("gorm:before_update").Register(spanBegin, cp.spanBegin("update"))
	_ = db.Callback().Update().After("gorm:after_update").Register(spanEnd, cp.spanEnd)

	_ = db.Callback().Raw().Before("gorm:raw").Register(spanBegin, cp.spanBegin("raw"))
	_ = db.Callback().Raw().After("gorm:raw").Register(spanEnd, cp.spanEnd)

	_ = db.Callback().Row().Before("gorm:row").Register(spanBegin, cp.spanBegin("row"))
	_ = db.Callback().Row().After("gorm:row").Register(spanEnd, cp.spanEnd)

	return nil
}

func (cp TracePlugin) spanBegin(processor string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		ctx, _ := cp.tracer.Start(db.Statement.Context, fmt.Sprintf("gorm.process.%s", processor))
		db.Statement.Context = ctx
	}
}

func (cp TracePlugin) spanEnd(db *gorm.DB) {
	span := trace.SpanFromContext(db.Statement.Context)
	defer span.End()

	err := db.Error
	var errMsg string
	if err != nil {
		errMsg = err.Error()
		span.RecordError(err)
		span.SetStatus(codes.Error, errMsg)
	}

	// TODO 隐藏私密参数
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	span.SetAttributes(
		attribute.String("gorm.sql", sql),
		attribute.String("gorm.error", errMsg),
		attribute.Int64("gorm.rows", db.Statement.RowsAffected),
	)
}
