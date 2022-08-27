package handler

import (
	"sync"
)

var (
	_registry Registry
	_lock     sync.Mutex
)

type Registry interface {
	Size() int
	Register(name string, handler TypeHandler)
	Handlers() []TypeHandler
}

func init() {
	_registry = initRegistry()
}

type registry struct {
	ctx map[string]TypeHandler
}

func (r *registry) Size() int {
	return len(r.ctx)
}

func (r *registry) Register(name string, handler TypeHandler) {
	r.ctx[name] = handler
}

func (r *registry) Handlers() []TypeHandler {
	_handlers := make([]TypeHandler, 0, _registry.Size())

	for _, handler := range r.ctx {
		_handlers = append(_handlers, handler)
	}

	return _handlers
}

func Register(name string, handler TypeHandler) {
	initRegistry().Register(name, handler)
}

func initRegistry() Registry {
	if _registry == nil {
		_lock.Lock()
		defer _lock.Unlock()
		if _registry == nil {
			_registry = &registry{
				ctx: make(map[string]TypeHandler, 0),
			}
		}
	}

	return _registry
}

func TypeHandlers() []TypeHandler {
	return _registry.Handlers()
}
