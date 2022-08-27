package handler

import (
	"fmt"
	"strings"

	"github.com/photowey/tsc/internal/data/keywords"
	"github.com/photowey/tsc/internal/data/types"
)

var _ TypeHandler = (*StringTypeHandler)(nil)

func init() {
	Register("string", &StringTypeHandler{})
}

type StringTypeHandler struct{}

func (h StringTypeHandler) Supports(v any) bool {
	_, ok := v.(types.String)

	return ok
}

func (h StringTypeHandler) Handle(k string, v types.Any) string {
	// readonly property?: string
	template := "%s%s%s: string // %s"
	readonlySymbol := ""
	if strings.Contains(k, keywords.ReadOnly) {
		readonlySymbol = keywords.ReadOnly + " "
	}
	requiredSymbol := "?"
	if strings.Contains(k, keywords.Required) {
		requiredSymbol = ""
	}

	return fmt.Sprintf(template, readonlySymbol, strings.Split(k, "|")[0], requiredSymbol, v)
}
