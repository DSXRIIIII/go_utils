package functionB

import (
	"context"
	"github.com/DSXRIIIII/go-utils/go-jaeger/tracer"
	"github.com/sirupsen/logrus"
)

func ChainEnd(ctx context.Context) string {
	logrus.Info("function End")
	return tracer.TraceID(ctx)
}
