package handlers

import (
	"fmt"

	"github.com/zalgonoise/logx/attr"
	"github.com/zalgonoise/logx/level"
	"github.com/zalgonoise/logx/records"
)

type multiHandler struct {
	handlers []Handler
}

// Multi will take any number of Handlers and return a multiHandler
// that batches the method calls similarly across all Handlers
//
// The resulting multiHandler is immutable
func Multi(h ...Handler) Handler {
	switch len(h) {
	case 0:
		return nil
	case 1:
		return h[0]
	default:
		var handlers []Handler
		for _, handler := range h {
			if handler == nil {
				continue
			}
			if mh, ok := handler.(multiHandler); ok {
				handlers = append(handlers, mh.handlers...)
				continue
			}
			handlers = append(handlers, handler)
		}
		return multiHandler{
			handlers: handlers,
		}
	}
}

// Enabled returns a boolean on whether the Handler is accepting
// records with log level `level`
func (mh multiHandler) Enabled(level level.Level) bool {
	for _, h := range mh.handlers {
		if ok := h.Enabled(level); !ok {
			return false // first handler filtering this level returns false
		}
	}
	return true // all handlers accept this level
}

// Handle will process the input Record, returning an error if raised
func (mh multiHandler) Handle(r records.Record) error {
	var err error
	for _, h := range mh.handlers {
		handlerErr := h.Handle(r)
		if handlerErr != nil {
			if err == nil {
				err = handlerErr
				continue
			}
			err = fmt.Errorf("%v -- %w", handlerErr, err)
		}
	}
	return err
}

// With will spawn a copy of this Handler with the input attributes
// `attrs`
func (mh multiHandler) With(attrs ...attr.Attr) Handler {
	newHandlers := make([]Handler, len(mh.handlers), len(mh.handlers))
	for idx, h := range mh.handlers {
		newHandlers[idx] = h.With(attrs...)
	}
	return Multi(newHandlers...)
}

// WithSource will spawn a new copy of this Handler with the setting
// to add a source file+line reference to `addSource` boolean
func (mh multiHandler) WithSource(addSource bool) Handler {
	newHandlers := make([]Handler, len(mh.handlers), len(mh.handlers))
	for idx, h := range mh.handlers {
		newHandlers[idx] = h.WithSource(addSource)
	}
	return Multi(newHandlers...)
}

// WithLevel will spawn a copy of this Handler with the input level `level`
// as a verbosity filter
func (mh multiHandler) WithLevel(level level.Level) Handler {
	newHandlers := make([]Handler, len(mh.handlers), len(mh.handlers))
	for idx, h := range mh.handlers {
		newHandlers[idx] = h.WithLevel(level)
	}
	return Multi(newHandlers...)
}

// WithReplaceFn will spawn a copy of this Handler with the input attribute
// replace function `fn`
func (mh multiHandler) WithReplaceFn(fn func(a attr.Attr) attr.Attr) Handler {
	newHandlers := make([]Handler, len(mh.handlers), len(mh.handlers))
	for idx, h := range mh.handlers {
		newHandlers[idx] = h.WithReplaceFn(fn)
	}
	return Multi(newHandlers...)
}
