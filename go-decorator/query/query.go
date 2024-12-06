package query

import (
	"context"
	"fmt"
	"github.com/DSXRIIIII/go-utils/go-decorator/command"
	"github.com/sirupsen/logrus"
)

type QueryCommand struct {
}

type QueryResult struct {
	ResultID string
}

type QueryDecoratorHandle command.CommandHandler[QueryCommand, *QueryResult]

type queryDecoratorHandle struct {
}

func (q queryDecoratorHandle) Handler(ctx context.Context, cmd QueryCommand) (*QueryResult, error) {
	fmt.Println("queryDecoratorHandle success")
	return &QueryResult{ResultID: "123"}, nil
}

func NewQueryCommandHandle(logger *logrus.Entry) QueryDecoratorHandle {
	return command.ApplyCommandDecorators[QueryCommand, *QueryResult](queryDecoratorHandle{}, logger)
}
