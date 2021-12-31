package mediator

import (
	"context"
	"reflect"
)

func newHandler(f reflect.Value) handler {
	return func(ctx context.Context, msg interface{}) (interface{}, error) {
		response := f.Call([]reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(msg),
		})

		if len(response) == 0 {
			return nil, nil
		}

		var err error
		if errVal, hasError := response[len(response)-1].Interface().(error); hasError {
			err = errVal
		}

		return response[0].Interface(), err
	}
}

func typeOf(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}
