package handlers

import (
	"github.com/zalgonoise/attr"
	"github.com/zalgonoise/logx/level"
	"github.com/zalgonoise/logx/records"
)

type unimplemented struct{}

func (unimplemented) Enabled(level level.Level) bool {
	return false
}
func (unimplemented) Handle(records.Record) error {
	return nil
}
func (u unimplemented) With(attrs ...attr.Attr) Handler {
	return u
}
func (u unimplemented) WithSource(addSource bool) Handler {
	return u
}
func (u unimplemented) WithLevel(level level.Level) Handler {
	return u
}
func (u unimplemented) WithReplaceFn(fn func(a attr.Attr) attr.Attr) Handler {
	return u
}

func Unimpl() Handler {
	return unimplemented{}
}
