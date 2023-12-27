package veun

import (
	"context"
)

type Error struct {
	Err error

	ctx context.Context
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) Context() context.Context {
	if e.ctx != nil {
		return e.ctx
	}
	return context.Background()
}

func (e *Error) WithContext(ctx context.Context) *Error {
	if ctx == nil {
		panic("nil context")
	}
	e2 := new(Error)
	*e2 = *e
	e2.ctx = ctx
	return e2
}
