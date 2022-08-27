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
	"fmt"
	"strings"

	"github.com/photowey/tsc/internal/data/keywords"
	"github.com/photowey/tsc/internal/data/types"
	"github.com/photowey/tsc/internal/funcs"
)

var _ TypeHandler = (*AnyMapTypeHandler)(nil)

func init() {
	Register("object", &AnyMapTypeHandler{})
}

type AnyMapTypeHandler struct{}

func (h AnyMapTypeHandler) Supports(v any) bool {
	_, ok := v.(types.AnyMap)

	return ok
}

func (h AnyMapTypeHandler) Handle(k string, v types.Any) string {
	pascalProperty := funcs.Pascal(strings.Split(k, "|")[0])
	template := "%s%s%s: %s // %s"
	readonlySymbol := ""
	if strings.Contains(k, keywords.ReadOnly) {
		readonlySymbol = keywords.ReadOnly + " "
	}
	requiredSymbol := "?"
	if strings.Contains(k, keywords.Required) {
		requiredSymbol = ""
	}

	return fmt.Sprintf(template, readonlySymbol, pascalProperty, requiredSymbol, pascalProperty, pascalProperty)
}
