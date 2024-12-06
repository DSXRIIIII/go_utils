package decorator

import (
	"context"
	"fmt"
	"github.com/DSXRIIIII/go-utils/go-decorator/command"
	"github.com/sirupsen/logrus"
)

type LoggingDecorator[C, R any] struct {
	Logger  *logrus.Entry
	Command command.CommandHandler[C, R]
}

func (d LoggingDecorator[C, R]) Handler(ctx context.Context, cmd C) (result R, err error) {
	logger := d.Logger.WithFields(logrus.Fields{
		"COMMADN":      cmd,
		"COMMADN_BODY": fmt.Sprintf("%#v", cmd),
	})
	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("success")
		} else {
			logger.Error("fail")
		}
	}()
	return d.Command.Handler(ctx, cmd)
}
