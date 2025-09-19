package logger

import "log/slog"

type slogAdapter struct {
	inner *slog.Logger
}

func (a slogAdapter) Debug(s string, fields ...Field) {
	a.inner.Debug(s, fieldsToAttr(fields)...)
}

func (a slogAdapter) Info(s string, fields ...Field) {
	a.inner.Info(s, fieldsToAttr(fields)...)
}

func (a slogAdapter) Error(s string, fields ...Field) {
	a.inner.Error(s, fieldsToAttr(fields)...)
}

func newSlogAdapter(inner *slog.Logger) *slogAdapter {
	return &slogAdapter{inner: inner}
}

func fieldsToAttr(fields []Field) []any {
	attrs := make([]any, 0, len(fields)*2)

	for _, f := range fields {
		attrs = append(attrs, f.Key, f.Value)
	}

	return attrs
}
