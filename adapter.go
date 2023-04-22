package logrus

import (
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/ydb-platform/ydb-go-sdk/v3/log"
)

var _ log.Logger = adapter{}

type adapter struct {
	l *logrus.Logger
}

func (a adapter) Log(params log.Params, msg string, fields ...log.Field) {
	a.l.WithFields(fieldsToFields(fields, params.Namespace)).Log(lvl2lvl(params.Level), msg)
}

func lvl2lvl(lvl log.Level) logrus.Level {
	switch lvl {
	case log.TRACE:
		return logrus.TraceLevel
	case log.DEBUG:
		return logrus.DebugLevel
	case log.INFO:
		return logrus.InfoLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	case log.FATAL:
		return logrus.FatalLevel
	default:
		return logrus.PanicLevel
	}
}

func fieldToField(field log.Field) (key string, value interface{}) {
	switch field.Type() {
	case log.IntType:
		return field.Key(), field.IntValue()
	case log.Int64Type:
		return field.Key(), field.Int64Value()
	case log.StringType:
		return field.Key(), field.StringValue()
	case log.BoolType:
		return field.Key(), field.BoolValue()
	case log.DurationType:
		return field.Key(), field.DurationValue()
	case log.StringsType:
		return field.Key(), field.StringsValue()
	case log.ErrorType:
		return "error", field.ErrorValue()
	case log.StringerType:
		return field.Key(), field.Stringer()
	default:
		return field.Key(), field.AnyValue()
	}
}

func fieldsToFields(fields []log.Field, namespace []string) (ff logrus.Fields) {
	ff = make(logrus.Fields, len(fields)+1)
	ff["scope"] = strings.Join(namespace, ".")
	for _, f := range fields {
		k, v := fieldToField(f)
		ff[k] = v
	}
	return ff
}
