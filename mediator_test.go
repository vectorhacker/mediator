package mediator_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vectorhacker/mediator"
)

type handler struct{}

type Message struct {
	Foo string
}

func (h handler) Handle(ctx context.Context, msg *Message) error {

	return nil
}

type invalidHandler struct{}

func (h invalidHandler) Handle() error {

	return nil
}

type takesNoContext struct{}

func (h takesNoContext) Handle(msg Message, foo string) error {

	return nil
}

func TestHandler(t *testing.T) {

	m, err := mediator.New(mediator.WithHandler(handler{}))
	assert.Nil(t, err)
	assert.NotNil(t, m)

	_, err = m.Send(context.Background(), &Message{"bar"})
	assert.Nil(t, err)
}

func TestHandlerFunc(t *testing.T) {
	m, err := mediator.New(mediator.WithHandlerFunc(func(ctx context.Context, msg Message) error {
		return nil
	}))

	assert.Nil(t, err)
	assert.NotNil(t, m)

	_, err = m.Send(context.Background(), Message{"bar"})
	assert.Nil(t, err)
}
func TestNoOptionsSent(t *testing.T) {
	_, err := mediator.New()
	assert.NotNil(t, err)
}

func TestHandlerNotFound(t *testing.T) {
	m, err := mediator.New(mediator.WithHandlerFunc(func(ctx context.Context, msg Message) error {
		return nil
	}))

	assert.Nil(t, err)
	assert.NotNil(t, m)

	_, err = m.Send(context.Background(), "yes")
	assert.NotNil(t, err)
}

func TestInvalidHandler(t *testing.T) {
	m, err := mediator.New(mediator.WithHandlerFunc(func(msg Message) error {
		return nil
	}))

	assert.NotNil(t, err)
	assert.Nil(t, m)

	m, err = mediator.New(mediator.WithHandlerFunc(func(msg Message, foo string) error {
		return nil
	}))

	assert.NotNil(t, err)
	assert.Nil(t, m)

	m, err = mediator.New(mediator.WithHandler(invalidHandler{}))
	assert.NotNil(t, err)
	assert.Nil(t, m)

	m, err = mediator.New(mediator.WithHandler(takesNoContext{}))
	assert.NotNil(t, err)
	assert.Nil(t, m)

	m, err = mediator.New(mediator.WithHandler(struct{}{}))
	assert.NotNil(t, err)
	assert.Nil(t, m)

	m, err = mediator.New(mediator.WithHandlerFunc(struct{}{}))
	assert.NotNil(t, err)
	assert.Nil(t, m)
}

func TestNilResponseWhenNoReturn(t *testing.T) {
	m, err := mediator.New(mediator.WithHandlerFunc(func(ctx context.Context, m Message) {

	}))

	assert.NotNil(t, m)
	assert.Nil(t, err)

	a, b := m.Send(context.Background(), Message{})
	assert.Nil(t, a)
	assert.Nil(t, b)
}

func TestReturnsErroValue(t *testing.T) {
	m, _ := mediator.New(mediator.WithHandlerFunc(func(ctx context.Context, m Message) error {
		return fmt.Errorf("sample")
	}))

	_, err := m.Send(context.Background(), Message{})
	assert.NotNil(t, err)
}
