package core

import (
	"context"
	"net/http"
	"time"

	"configcenter/src/common/errors"
	"configcenter/src/common/language"
)

// 定义获取上下文参数的数据结构
type ContextParams struct {
	context.Context
	Header          http.Header
	SupplierAccount string
	User            string
	ReqID           string
	Error           errors.DefaultCCErrorIf
	Lang            language.DefaultCCLanguageIf
}

// Deadline overwrite Context Deadline methods
func (c ContextParams) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

// Done overwrite Context Done methods
func (c ContextParams) Done() <-chan struct{} {
	return c.Context.Done()
}

// Err overwrite Context Err methods
func (c ContextParams) Err() error {
	return c.Context.Err()
}

// Value overwrite Context Value methods
func (c ContextParams) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
