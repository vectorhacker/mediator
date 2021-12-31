package mediator

import (
	"context"
	"fmt"
	"reflect"
)

type handler func(ctx context.Context, msg interface{}) (interface{}, error)

// Mediator is a simple implementation of the mediator pattern
type Mediator struct {
	handlers map[reflect.Type]handler
}

// New creates a new mediator with the options set
func New(options ...Option) (*Mediator, error) {

	m := &Mediator{
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

// Send delegates a message to the appropriate handler
func (m Mediator) Send(
	ctx context.Context,
	message interface{},
) (response interface{}, err error) {

	h, ok := m.handlers[typeOf(message)]
	if !ok {
		return nil, ErrHandlerNotFound
	}

	return h(ctx, message)
}
