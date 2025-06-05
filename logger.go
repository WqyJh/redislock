package redislock

import (
	"context"
	"log/slog"
)

type discardHandler struct{}

func (discardHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (discardHandler) Handle(context.Context, slog.Record) error { return nil }
func (discardHandler) WithAttrs(attrs []slog.Attr) slog.Handler  { return discardHandler{} }
func (discardHandler) WithGroup(name string) slog.Handler        { return discardHandler{} }

func DiscardLogger() *slog.Logger {
	return slog.New(discardHandler{})
}
