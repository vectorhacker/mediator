# mediator

Simple mediator implementation for Go


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