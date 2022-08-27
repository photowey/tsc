package handler

import (
	"fmt"
	"strings"

	"github.com/photowey/tsc/internal/data/keywords"
	"github.com/photowey/tsc/internal/data/types"
)

var _ TypeHandler = (*BooleanTypeHandler)(nil)

func init() {
	Register("boolean", &BooleanTypeHandler{})
}

type BooleanTypeHandler struct{}

func (h BooleanTypeHandler) Supports(v any) bool {
	_, ok := v.(types.Boolean)

	return ok
}

func (h BooleanTypeHandler) Handle(k string, v types.Any) string {
	// readonly property?: boolean
	template := "%s%s%s: boolean // %t"
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
