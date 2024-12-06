package command

import (
	"context"
	"github.com/DSXRIIIII/go-utils/go-decorator/decorator"
	"github.com/sirupsen/logrus"
)

type CommandHandler[C, R any] interface {
	Handler(ctx context.Context, cmd C) (R, error)
}

func ApplyCommandDecorators[C, R any](handle CommandHandler[C, R], logger *logrus.Entry) CommandHandler[C, R] {
	return decorator.LoggingDecorator[C, R]{
		Logger:  logger,
		Command: handle,
	}
}
