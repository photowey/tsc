/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
