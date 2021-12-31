package mediator

import (
	"context"
	"fmt"
	"reflect"
)

type handler func(ctx context.Context, msg interface{}) (interface{}, error)

// Mediator is a simple implementation of the mediator pattern
type Mediator interface {
	// Send delegates a message to the appropriate handler
	Send(ctx context.Context, msg interface{}) (res interface{}, err error)
}
type mediator struct {
	handlers map[reflect.Type]handler
}

// New creates a new mediator with the options set
func New(options ...Option) (Mediator, error) {

	m := &mediator{
		handlers: make(map[reflect.Type]handler),
	}

	if len(options) == 0 {
		return nil, fmt.Errorf("must set at least one option")
	}

	for _, opt := range options {
		if err := opt(m); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m mediator) Send(ctx context.Context, msg interface{}) (interface{}, error) {

	h, ok := m.handlers[typeOf(msg)]
	if !ok {
		return nil, ErrHandlerNotFound
	}

	return h(ctx, msg)
}
