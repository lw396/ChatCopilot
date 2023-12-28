package redis

import "go.opentelemetry.io/otel/trace"

type options struct {
	host     string
	port     int
	user     string
	password string
	db       int
	packer   Packer
	tracer   trace.Tracer
}

func defaultOptions() *options {
	return &options{
		host:   "localhost",
		port:   6379,
		db:     0,
		packer: JSONPacker,
	}
}

type Option func(o *options)

func WithAddress(host string, port int) Option {
	return func(o *options) {
		o.host = host
		o.port = port
	}
}

func WithAuth(user, password string) Option {
	return func(o *options) {
		o.user = user
		o.password = password
	}
}

func WithDB(db int) Option {
	return func(o *options) {
		o.db = db
	}
}

func WithPacker(packer Packer) Option {
	return func(o *options) {
		o.packer = packer
	}
}

func WithTracer(tracer trace.Tracer) Option {
	return func(o *options) {
		o.tracer = tracer
	}
}
