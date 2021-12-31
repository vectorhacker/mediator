# mediator

Simple mediator implementation for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/vectorhacker/mediator)](https://goreportcard.com/report/github.com/vectorhacker/mediator)

Example 

```go

type MessageHandler struct{}

type Message struct {
	Foo string
}

type Response struct {
    Result string
}

func (h MessageHandler) Handle(ctx context.Context, msg *Message) (Response, error) {

	return Response{ msg.Foo + " bar "}, nil
}


m, err := mediator.New(mediator.WithHandler(&MessageHandler{})) 

r, err := m.Send(context.Background(), &Message{ "foo" })
// ...

```