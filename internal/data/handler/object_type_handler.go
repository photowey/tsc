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
