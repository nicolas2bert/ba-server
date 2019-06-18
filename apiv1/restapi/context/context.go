package context

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type logkey string

var lk = logkey("log")

func AddLog(l *log.Entry, ctx context.Context) context.Context {
	return context.WithValue(ctx, lk, l)
}

func GetLog(ctx context.Context, handlerName string) *log.Entry {
	l := ctx.Value(lk).(*log.Entry)
	return l.WithField("handlerName", handlerName)
}
