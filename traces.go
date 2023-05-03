package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/log"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

func WithTraces(l *logrus.Logger, d trace.Detailer, opts ...log.Option) ydb.Option {
	a := adapter{l: l}
	return ydb.WithLogger(a, d, opts...)
}

func Table(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Table {
	return log.Table(&adapter{l: l}, d, opts...)
}

func Topic(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Topic {
	return log.Topic(&adapter{l: l}, d, opts...)
}

func Driver(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Driver {
	return log.Driver(&adapter{l: l}, d, opts...)
}

func Coordination(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Coordination {
	return log.Coordination(&adapter{l: l}, d, opts...)
}

func Discovery(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Discovery {
	return log.Discovery(&adapter{l: l}, d, opts...)
}

func Ratelimiter(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Ratelimiter {
	return log.Ratelimiter(&adapter{l: l}, d, opts...)
}

func Scheme(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Scheme {
	return log.Scheme(&adapter{l: l}, d, opts...)
}

func Scripting(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.Scripting {
	return log.Scripting(&adapter{l: l}, d, opts...)
}

func DatabaseSQL(l *logrus.Logger, d trace.Detailer, opts ...log.Option) trace.DatabaseSQL {
	return log.DatabaseSQL(&adapter{l: l}, d, opts...)
}
