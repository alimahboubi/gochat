package interfaces

import "context"

type Query interface {
	QueryType() string
}

type QueryHandler interface {
	Handle(ctx *context.Context, query Query)
}
