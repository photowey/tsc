package handler

import (
	"fmt"
	"strings"

	"github.com/photowey/tsc/internal/data/keywords"
	"github.com/photowey/tsc/internal/data/types"
)

var _ TypeHandler = (*NumberTypeHandler)(nil)

func init() {
	Register("number", &NumberTypeHandler{})
}

type NumberTypeHandler struct{}

func (h NumberTypeHandler) Supports(v any) bool {
	_, ok := v.(types.Number)

	return ok
}

func (h NumberTypeHandler) Handle(k string, v types.Any) string {
	// readonly property?: number
	template := "%s%s%s: number // %v"
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
