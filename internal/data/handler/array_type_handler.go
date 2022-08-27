package handler

import (
	"fmt"
	"strings"

	"github.com/photowey/tsc/internal/data/keywords"
	"github.com/photowey/tsc/internal/data/types"
	"github.com/photowey/tsc/internal/funcs"
)

var _ TypeHandler = (*ArrayTypeHandler)(nil)

func init() {
	Register("array", &ArrayTypeHandler{})
}

type ArrayTypeHandler struct{}

func (h ArrayTypeHandler) Supports(v any) bool {
	_, ok := v.([]any)

	return ok
}

func (h ArrayTypeHandler) Handle(k string, v types.Any) string {
	pascalProperty := funcs.Pascal(strings.Split(k, "|")[0])
	template := "%s%s%s: %s[] // %s"
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
