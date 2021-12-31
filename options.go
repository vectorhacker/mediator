package mediator

import (
	"context"
	"fmt"
	"reflect"
)

// Option is a function that modifies the mediator with certain options
type Option func(*mediator) error

// WithHandlerFunc adds a handler func which takes two arguments, a context
// and a message.
func WithHandlerFunc(f interface{}) Option {
	return func(m *mediator) error {
		t := typeOf(f)
		if t.Kind() != reflect.Func {
			return fmt.Errorf("handler func must be a function")
		}

		if t.NumIn() != 2 {
			return fmt.Errorf("handler must have two arguments")
		}

		contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
		firstArgument := t.In(0)
		secondArgument := t.In(1)
		handlerTakesContext := firstArgument.Implements(contextType)

		if !handlerTakesContext {
			return fmt.Errorf("handler must take context.Context as first argument")
		}

		h := reflect.ValueOf(f)
		m.handlers[secondArgument] = newHandler(h)

		return nil
	}
}

// WithHandler adds a message handler to the mediator
// a handler must have an Handle method with two arguments
// and two return values
func WithHandler(h interface{}) Option {
	return func(m *mediator) error {
		t := typeOf(h)
		method, ok := t.MethodByName("Handle")
		if !ok {
			return fmt.Errorf("handler must have method Handle")
		}

		if method.Type.NumIn() != 3 {
			// return an error handler
			return fmt.Errorf("handler must have method Handle with two arguments")
		}

		contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
		firstArgument := method.Type.In(1)
		secondArgument := method.Type.In(2)
		handlerTakesContext := firstArgument.Implements(contextType)

		if !handlerTakesContext {
			return fmt.Errorf("handle method must take context.Context as first argument")
		}

		if secondArgument.Kind() == reflect.Ptr {
			secondArgument = secondArgument.Elem()
		}

		m.handlers[secondArgument] = newHandler(reflect.ValueOf(h).MethodByName("Handle"))

		return nil
	}
}
