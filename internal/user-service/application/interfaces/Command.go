package interfaces

import "context"

type Command interface {
	CommandType() string
}

type CommandHandler[TCommand Command, TResult any] interface {
	Handle(ctx *context.Context, command TCommand) (TResult, error)
}
